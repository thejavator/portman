package tui

import (
	"time"

	"portman/config"
	"portman/system"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time
type portsLoadedMsg []system.PortInfo
type detailsLoadedMsg system.ProcessDetails

// Model represents the application state
type Model struct {
	width      int
	height     int
	ports      []system.PortInfo
	selected   int
	loading    bool
	appVersion string

	activeTab int
	searchBar textinput.Model
	config    config.AppConfig
	styles    Styles

	showDetails bool
	confirmKill bool
	details     system.ProcessDetails
}

// InitialModel initializes the state
func InitialModel(version string) Model {
	cfg := config.LoadConfig()
	
	ti := textinput.New()
	ti.Placeholder = "Search (port, name, pid)..."
	ti.CharLimit = 156
	ti.Width = 40

	return Model{
		loading:    true,
		activeTab:  0,
		searchBar:  ti,
		appVersion: version,
		config:    cfg,
		styles:    GetStyles(cfg.Theme),
	}
}

// Init lance les commandes initiales
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.fetchPorts(),
		textinput.Blink,
		tickCmd(),
	)
}

func (m Model) fetchPorts() tea.Cmd {
	return func() tea.Msg {
		ports, _ := system.ScanPorts(&m.config)
		return portsLoadedMsg(ports)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
