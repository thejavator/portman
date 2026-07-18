//go:build linux

package system

import (
	"os/exec"
)

// ToggleBlock blocks or unblocks an incoming/outgoing port via iptables
func ToggleBlock(port string) error {
	action := "-A" // Add
	if blockedPorts[port] {
		action = "-D" // Delete
		delete(blockedPorts, port)
	} else {
		blockedPorts[port] = true
	}

	// Block/Unblock TCP
	cmdTCP := exec.Command("iptables", action, "INPUT", "-p", "tcp", "--dport", port, "-j", "DROP")
	if err := cmdTCP.Run(); err != nil {
		return err
	}

	// Block/Unblock UDP
	cmdUDP := exec.Command("iptables", action, "INPUT", "-p", "udp", "--dport", port, "-j", "DROP")
	return cmdUDP.Run()
}
