package system

type ProcessInfo struct {
	ID          string
	Address     string
	Started     string
	IsFavorite  bool
	Conflict    bool
	Blocked     bool
}
