package system

var blockedPorts = make(map[string]bool)

// IsBlocked returns true if port is blocked
func IsBlocked(port string) bool {
	return blockedPorts[port]
}
