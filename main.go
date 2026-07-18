package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"portman/tui"

	tea "github.com/charmbracelet/bubbletea"
)
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		if version == "dev" {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					commit = setting.Value
					if len(commit) > 7 {
						commit = commit[:7]
					}
				}
				if setting.Key == "vcs.time" {
					date = setting.Value
				}
				if setting.Key == "vcs.modified" && setting.Value == "true" {
					commit += "-dirty"
				}
			}
		}
	}
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Show application version")
	flag.BoolVar(&showVersion, "v", false, "Show application version")
	flag.Parse()

	if showVersion {
		fmt.Printf("portman version %s, commit %s, built at %s\n", version, commit, date)
		os.Exit(0)
	}

	// Sudo check (optional for launching, but required for some actions)
	if os.Geteuid() != 0 {
		fmt.Println("⚠️  Warning: portman is not running with sudo.")
		fmt.Println("    Some actions (killing root processes, pfctl firewall) will fail.")
		fmt.Println("    Press Enter to continue or Ctrl+C to quit.")
		// We could block here, but we allow the user to continue for now
	}

	p := tea.NewProgram(tui.InitialModel(version), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fatal error: %v", err)
		os.Exit(1)
	}
}
