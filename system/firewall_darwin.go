//go:build darwin

package system

import (
	"fmt"
	"os/exec"
)

// ToggleBlock blocks or unblocks an incoming/outgoing port via pfctl
func ToggleBlock(port string) error {
	if blockedPorts[port] {
		delete(blockedPorts, port)
	} else {
		blockedPorts[port] = true
	}

	// Active pfctl s'il ne l'est pas
	exec.Command("pfctl", "-e").Run()

	var rules string
	for p := range blockedPorts {
		rules += fmt.Sprintf("block drop in proto tcp from any to any port %s\n", p)
		rules += fmt.Sprintf("block drop in proto udp from any to any port %s\n", p)
	}

	if len(blockedPorts) == 0 {
		return exec.Command("pfctl", "-a", "com.portman.block", "-F", "rules").Run()
	}

	cmd := exec.Command("pfctl", "-a", "com.portman.block", "-f", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	
	go func() {
		defer stdin.Close()
		fmt.Fprint(stdin, rules)
	}()

	return cmd.Run()
}
