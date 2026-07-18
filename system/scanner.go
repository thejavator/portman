package system

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"portman/config"

	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

// Category de port
type Category string

const (
	CatSystem   Category = "System"
	CatApp      Category = "App"
	CatDev      Category = "Dev"
	CatNetwork  Category = "Network"
	CatOther    Category = "Other"
)

// ProcessInfo is in process.go, we use PortInfo
type PortInfo struct {
	PID         string
	ProcessName string
	Protocol    string
	Port        string
	Category    Category
	Address     string
	Started     string
	IsFavorite  bool
	Conflict    bool
	Blocked     bool
}

type ProcessDetails struct {
	PortInfo      PortInfo
	Path          string
	FullCommand   string
	Cwd           string
	User          string
	Started       string
	PPID          string
	ParentProcess string
	URL           string
}

var devApps = []string{"node", "python", "go", "ruby", "php", "java", "npm", "yarn", "bun", "docker", "nginx", "httpd"}
var sysApps = []string{"launchd", "configd", "mDNSResponder", "syslogd", "cupsd", "syspolicyd", "rapportd", "coreservicesd", "svchost", "System"}
var netApps = []string{"vpnagentd", "acumbrellaagent", "com.cisco", "wireguard", "tailscaled", "openvpn"}
var appApps = []string{"ollama"}

func isExactMatch(cmdLower, app string) bool {
	if cmdLower == app {
		return true
	}
	if strings.HasSuffix(cmdLower, "/"+app) || strings.HasSuffix(cmdLower, "\\"+app) || strings.HasSuffix(cmdLower, app+".exe") {
		return true
	}
	if strings.HasPrefix(cmdLower, app+" ") {
		return true
	}
	return false
}

func determineCategory(cmd string) Category {
	cmdLower := strings.ToLower(cmd)
	for _, app := range devApps {
		if isExactMatch(cmdLower, app) {
			return CatDev
		}
	}
	for _, app := range netApps {
		if isExactMatch(cmdLower, app) {
			return CatNetwork
		}
	}
	for _, app := range sysApps {
		if isExactMatch(cmdLower, app) {
			return CatSystem
		}
	}
	for _, app := range appApps {
		if isExactMatch(cmdLower, app) {
			return CatApp
		}
	}
	if len(cmd) > 0 && cmd[0] >= 'A' && cmd[0] <= 'Z' {
		return CatApp
	}
	return CatOther
}

func ScanPorts(cfg *config.AppConfig) ([]PortInfo, error) {
	conns, err := net.Connections("inet")
	if err != nil {
		return nil, err
	}

	var ports []PortInfo
	seen := make(map[string]bool)

	for _, conn := range conns {
		if conn.Pid == 0 {
			continue // usually system idle process or unknown
		}

		proto := "TCP"
		if conn.Type == 2 { // SOCK_DGRAM
			proto = "UDP"
		} else if conn.Status != "LISTEN" && conn.Status != "LISTEN " {
			continue // Only interested in listening ports for TCP
		}

		portStr := strconv.Itoa(int(conn.Laddr.Port))
		addr := conn.Laddr.IP

		key := fmt.Sprintf("%d:%s:%s", conn.Pid, proto, portStr)
		if seen[key] {
			continue
		}
		seen[key] = true

		proc, err := process.NewProcess(conn.Pid)
		cmdName := "unknown"
		started := ""
		if err == nil {
			if name, err := proc.Name(); err == nil {
				cmdName = name
			}
			if ct, err := proc.CreateTime(); err == nil {
				started = time.UnixMilli(ct).Format("Mon Jan 2 15:04:05 2006")
			}
		}

		ports = append(ports, PortInfo{
			PID:         strconv.Itoa(int(conn.Pid)),
			ProcessName: cmdName,
			Protocol:    proto,
			Port:        portStr,
			Category:    determineCategory(cmdName),
			Address:     addr,
			Started:     started,
			IsFavorite:  cfg.IsFavorite(portStr),
			Blocked:     IsBlocked(portStr),
		})
	}

	// Detect conflicts
	portCount := make(map[string]map[string]bool)
	for _, p := range ports {
		key := p.Protocol + ":" + p.Port
		if portCount[key] == nil {
			portCount[key] = make(map[string]bool)
		}
		portCount[key][p.PID] = true
	}

	for i := range ports {
		key := ports[i].Protocol + ":" + ports[i].Port
		if len(portCount[key]) > 1 {
			ports[i].Conflict = true
		}
	}

	return ports, nil
}

func GetProcessDetails(p PortInfo) ProcessDetails {
	details := ProcessDetails{
		PortInfo: p,
	}

	pid, err := strconv.ParseInt(p.PID, 10, 32)
	if err != nil {
		return details
	}

	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		return details
	}

	if ppid, err := proc.Ppid(); err == nil {
		details.PPID = strconv.Itoa(int(ppid))
		if ppid != 0 {
			if pProc, err := process.NewProcess(ppid); err == nil {
				if pName, err := pProc.Name(); err == nil {
					details.ParentProcess = fmt.Sprintf("%s (%d)", pName, ppid)
				}
			}
		}
	}

	if user, err := proc.Username(); err == nil {
		details.User = user
	}

	if cwd, err := proc.Cwd(); err == nil {
		details.Cwd = cwd
	}

	if cmdLine, err := proc.Cmdline(); err == nil {
		details.FullCommand = cmdLine
	}

	if exe, err := proc.Exe(); err == nil {
		details.Path = exe
	} else if len(details.FullCommand) > 0 {
		details.Path = strings.Split(details.FullCommand, " ")[0]
	}

	if p.Protocol == "TCP" && p.Address != "0.0.0.0" && p.Address != "::" && p.Address != "*" && p.Address != "" {
		host := "localhost"
		if !strings.Contains(p.Address, "127.0.0.1") && !strings.Contains(p.Address, "::1") {
			host = p.Address
		}
		details.URL = "http://" + host + ":" + p.Port
	} else if p.Protocol == "TCP" {
		details.URL = "http://localhost:" + p.Port
	}

	return details
}
