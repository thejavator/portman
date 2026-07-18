//go:build darwin

package system

import (
	"fmt"
	"os/exec"
)

// KillProcess sends a signal (15 or 9) to a PID
func KillProcess(pid string, force bool) error {
	sig := "-15"
	if force {
		sig = "-9"
	}
	cmd := exec.Command("kill", sig, pid)
	return cmd.Run()
}

// OpenBrowser opens the URL in the default browser
func OpenBrowser(port string) error {
	url := fmt.Sprintf("http://localhost:%s", port)
	cmd := exec.Command("open", url)
	return cmd.Run()
}
