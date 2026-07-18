package tui

import (
	"fmt"
	"os/exec"
	"strings"

	"portman/system"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "y":
			if m.confirmKill {
				filtered := m.getFilteredPorts()
				if len(filtered) > 0 && m.selected < len(filtered) {
					system.KillProcess(filtered[m.selected].PID, false)
					m.confirmKill = false
					return m, m.fetchPorts()
				}
			}
		case "n":
			if m.confirmKill {
				m.confirmKill = false
				return m, nil
			}
		case "ctrl+r":
			m.searchBar.Focus()
			return m, nil
		case "tab":
			m.activeTab = (m.activeTab + 1) % 7
		case "shift+tab":
			m.activeTab--
			if m.activeTab < 0 {
				m.activeTab = 6
			}
		case "q", "esc":
			if m.showDetails {
				m.showDetails = false
				return m, nil
			}
			if msg.String() == "esc" {
				if m.confirmKill {
					m.confirmKill = false
					return m, nil
				}
				if m.searchBar.Focused() {
					m.searchBar.Blur()
				} else {
					m.searchBar.SetValue("")
				}
			} else if !m.searchBar.Focused() && !m.confirmKill {
				return m, tea.Quit
			}
		case "enter":
			filtered := m.getFilteredPorts()
			if !m.showDetails && len(filtered) > 0 && m.selected < len(filtered) {
				port := filtered[m.selected]
				return m, func() tea.Msg {
					return detailsLoadedMsg(system.GetProcessDetails(port))
				}
			}
		case "up":
			if m.showDetails {
				return m, nil
			}
			if m.selected > 0 {
				m.selected--
			}
		case "k":
			if m.showDetails {
				return m, nil
			}
			if !m.searchBar.Focused() && m.selected > 0 {
				m.selected--
			}
		case "down":
			if m.showDetails {
				return m, nil
			}
			if m.selected < len(m.getFilteredPorts())-1 {
				m.selected++
			}
		case "j":
			if m.showDetails {
				return m, nil
			}
			if !m.searchBar.Focused() && m.selected < len(m.getFilteredPorts())-1 {
				m.selected++
			}
		case "*":
			filtered := m.getFilteredPorts()
			if len(filtered) > 0 && m.selected < len(filtered) {
				p := filtered[m.selected]
				m.config.ToggleFavorite(p.Port)
				return m, m.fetchPorts()
			}
		case "K":
			if !m.searchBar.Focused() && !m.showDetails {
				m.confirmKill = true
				return m, nil
			}
		case "o":
			if !m.searchBar.Focused() && !m.showDetails {
				filtered := m.getFilteredPorts()
				if len(filtered) > 0 && m.selected < len(filtered) {
					p := filtered[m.selected]
					if p.Protocol == "TCP" {
						exec.Command("open", fmt.Sprintf("http://localhost:%s", p.Port)).Start()
					}
				}
			}
		case "b", "ctrl+b":
			if msg.String() == "ctrl+b" || (!m.searchBar.Focused() && !m.confirmKill) {
				filtered := m.getFilteredPorts()
				if len(filtered) > 0 && m.selected < len(filtered) {
					p := filtered[m.selected]
					m.config.ToggleFavorite(p.Port)
					return m, m.fetchPorts()
				}
			}
		}

		var cmd tea.Cmd
		m.searchBar, cmd = m.searchBar.Update(msg)
		cmds = append(cmds, cmd)

	case portsLoadedMsg:
		m.ports = msg
		m.loading = false
		filtered := m.getFilteredPorts()
		if m.selected >= len(filtered) && len(filtered) > 0 {
			m.selected = len(filtered) - 1
		}

	case detailsLoadedMsg:
		m.details = system.ProcessDetails(msg)
		m.showDetails = true

	case tickMsg:
		if !m.showDetails {
			cmds = append(cmds, m.fetchPorts(), tickCmd())
		} else {
			cmds = append(cmds, tickCmd())
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) getFilteredPorts() []system.PortInfo {
	var res []system.PortInfo
	search := m.searchBar.Value()

	for _, p := range m.ports {
		// Tabs
		if m.activeTab == 1 && p.Category != system.CatSystem {
			continue
		}
		if m.activeTab == 2 && p.Category != system.CatApp {
			continue
		}
		if m.activeTab == 3 && p.Category != system.CatNetwork {
			continue
		}
		if m.activeTab == 4 && p.Category != system.CatDev {
			continue
		}
		if m.activeTab == 5 && p.Category != system.CatOther {
			continue
		}
		if m.activeTab == 6 && !p.IsFavorite {
			continue
		}

		// Search
		if search != "" {
			match := strings.Contains(strings.ToLower(p.ProcessName), strings.ToLower(search)) ||
				strings.Contains(p.Port, search) ||
				strings.Contains(strings.ToLower(p.PID), strings.ToLower(search))
			if !match {
				continue
			}
		}

		res = append(res, p)
	}
	return res
}
