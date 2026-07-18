package main

import (
	"flag"
	"fmt"
	"os"

	"portman/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Show application version")
	flag.BoolVar(&showVersion, "v", false, "Show application version")
	flag.Parse()

	if showVersion {
		fmt.Println("portman version v0.1.0")
		os.Exit(0)
	}

	// Sudo check (optional for launching, but required for some actions)
	if os.Geteuid() != 0 {
		fmt.Println("⚠️  Warning: portman is not running with sudo.")
		fmt.Println("    Some actions (killing root processes, pfctl firewall) will fail.")
		fmt.Println("    Press Enter to continue or Ctrl+C to quit.")
		// We could block here, but we allow the user to continue for now
	}

	p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fatal error: %v", err)
		os.Exit(1)
	}
}
