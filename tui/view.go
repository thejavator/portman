package tui

import (
	"fmt"
	"strings"

	"portman/system"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderCategory(cat system.Category) string {
	switch cat {
	case system.CatApp:
		return m.styles.CategoryApp.Render("APP ")
	case system.CatSystem:
		return m.styles.CategorySys.Render("SYS ")
	case system.CatDev:
		return m.styles.CategoryDev.Render("DEV ")
	case system.CatNetwork:
		return m.styles.CategoryNet.Render("NET ")
	default:
		return m.styles.CategoryOth.Render("OTH ")
	}
}

// View generates the text interface
func (m Model) View() string {
	var content string
	if m.showDetails {
		content = m.detailsView()
	} else {
		// --- 1. HEADER ---
		headerBuilder := strings.Builder{}
		titleArt := `
 ___  ___  ___  _____ __  __  ___  _  _ 
| _ \/ _ \| _ \|_   _|  \/  |/ _ \| \| |
|  _/ (_) |   /  | | | |\/| |  _  | .  |
|_|  \___/|_|_\  |_| |_|  |_|_| |_|_|\_|`
		headerBuilder.WriteString(m.styles.Title.Render(strings.TrimPrefix(titleArt, "\n")) + "\n\n")
		headerBuilder.WriteString(m.styles.Description.Render("Advanced port manager ") + m.styles.Version.Render("v0.1.0") + "\n")
		headerStr := headerBuilder.String()

		// --- 2. TABS ---
		counts := make([]int, 7)
		for _, p := range m.ports {
			counts[0]++
			switch p.Category {
			case system.CatSystem:
				counts[1]++
			case system.CatApp:
				counts[2]++
			case system.CatNetwork:
				counts[3]++
			case system.CatDev:
				counts[4]++
			case system.CatOther:
				counts[5]++
			}
			if p.IsFavorite {
				counts[6]++
			}
		}

		tabs := []string{"All", "System", "Apps", "Network", "Dev", "Other", "Favorites"}
		var renderedTabs []string
		for i, t := range tabs {
			label := fmt.Sprintf("%s (%d)", t, counts[i])
			if m.activeTab == i {
				renderedTabs = append(renderedTabs, m.styles.ActiveTab.Render(label))
			} else {
				renderedTabs = append(renderedTabs, m.styles.InactiveTab.Render(label))
			}
		}
		tabsContent := strings.Join(renderedTabs, m.styles.Subtle.Render(" │ "))
		tabsStr := m.styles.Panel.Render(tabsContent)

		// --- 3. SEARCH ---
		searchStr := "🔍 Search: " + m.searchBar.View()
		if m.searchBar.Focused() {
			searchStr = m.styles.ActivePanel.Width(m.width - 2).Render(searchStr)
		} else {
			searchStr = m.styles.Panel.Width(m.width - 2).Render(searchStr)
		}

		// --- 4. TABLE ---
		tableBuilder := strings.Builder{}
		filtered := m.getFilteredPorts()

		if len(filtered) == 0 {
			if m.loading {
				tableBuilder.WriteString("Loading...\n")
			} else {
				tableBuilder.WriteString("No ports found.\n")
			}
		} else {
			headerRow := fmt.Sprintf("  %-8s %-6s %-25s %-6s %-10s %-18s %-15s %s", "PORT", "PROTO", "COMMAND", "PID", "TYPE", "ADDRESS", "STARTED", "FAV")
			tableBuilder.WriteString(m.styles.Header.Render(headerRow) + "\n")
			tableBuilder.WriteString(m.styles.Subtle.Render(strings.Repeat("─", m.width-4)) + "\n")

			// Viewport calculation for scrolling
			availableHeight := m.height - 24
			if availableHeight < 1 {
				availableHeight = 1
			}

			start := 0
			end := len(filtered)

			if len(filtered) > availableHeight {
				start = m.selected - availableHeight/2
				if start < 0 {
					start = 0
				}
				end = start + availableHeight
				if end > len(filtered) {
					end = len(filtered)
					start = end - availableHeight
				}
			}

			// Rows
			for i := start; i < end; i++ {
				port := filtered[i]
				cursor := "  "
				rowStyle := m.styles.NormalRow
				if port.Conflict {
					rowStyle = m.styles.Error
				}
				if m.selected == i {
					cursor = "> "
					rowStyle = m.styles.SelectedRow
					if port.Conflict {
						rowStyle = m.styles.Error.Copy().Bold(true)
					}
				}

				catStr := m.renderCategory(port.Category)
				catStrPadded := lipgloss.NewStyle().Width(10).Render(catStr)

				favStr := " "
				if port.IsFavorite {
					favStr = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true).Render("★")
				}
				favStr += " "

				pname := port.ProcessName
				if len(pname) > 24 {
					pname = pname[:21] + "..."
				}
				addr := port.Address
				if len(addr) > 17 {
					addr = addr[:14] + "..."
				}
				started := port.Started
				if len(started) > 14 {
					started = started[:14]
				}

				row := fmt.Sprintf("%s%-8s %-6s %-25s %-6s %s %-18s %-15s %s",
					cursor, port.Port, port.Protocol, pname, port.PID, catStrPadded, addr, started, favStr)

				tableBuilder.WriteString(rowStyle.Render(row) + "\n")
			}
		}

		tableStr := m.styles.Panel.Width(m.width - 2).Render(tableBuilder.String())

		// --- 5. FOOTER ---
		footerBuilder := strings.Builder{}
		if m.confirmKill {
			if len(filtered) > 0 && m.selected < len(filtered) {
				p := filtered[m.selected]
				confirmMsg := fmt.Sprintf("⚠️  Are you sure you want to kill %s (PID: %s)? [y/n]", p.ProcessName, p.PID)
				footerBuilder.WriteString(m.styles.Error.Copy().Bold(true).Render(confirmMsg))
			}
		} else {
			footerBuilder.WriteString(m.styles.Subtle.Render("Nav: [↑/↓] | Search: [Ctrl+R] | Tabs: [Tab] | Fav: [b/Ctrl+B] | Kill: [K] | Open: [o] | Quit: [Ctrl+C]"))
		}
		footerStr := footerBuilder.String()

		content = lipgloss.JoinVertical(lipgloss.Left, headerStr, tabsStr, searchStr, tableStr, footerStr)
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Left, lipgloss.Top, content)
}

func (m Model) detailsView() string {
	d := m.details
	p := d.PortInfo

	// --- Header Box ---
	portBlock := lipgloss.NewStyle().Width(15).Render(fmt.Sprintf("  %s %s", m.styles.SelectedRow.Render(p.Port), m.styles.Subtle.Render(p.Protocol)))
	nameBlock := lipgloss.NewStyle().Bold(true).Width(35).Render(p.ProcessName)
	parentBlock := m.styles.Subtle.Copy().Width(25).Render(d.ParentProcess)
	pidBlock := m.styles.Subtle.Copy().Width(10).Render(p.PID)
	catBlock := m.renderCategory(p.Category)

	headerContent := lipgloss.JoinHorizontal(lipgloss.Left, portBlock, nameBlock, parentBlock, pidBlock, catBlock)
	headerBox := m.styles.Panel.Width(m.width - 2).Render(headerContent)

	// --- Details Box ---
	detailsBuilder := strings.Builder{}
	labelStyle := m.styles.Subtle.Copy().Width(20)
	valueStyle := m.styles.NormalRow

	addRow := func(label, value string) {
		if value != "" {
			detailsBuilder.WriteString(fmt.Sprintf("  %s %s\n\n", labelStyle.Render(label), valueStyle.Render(value)))
		}
	}

	addRow("Process Name", p.ProcessName)
	addRow("Path", d.Path)
	addRow("Full command", d.FullCommand)
	addRow("Working directory", d.Cwd)
	addRow("User", d.User)
	addRow("Started time", d.Started)
	addRow("Parent process", d.ParentProcess)

	addRow("Address", p.Address)

	if d.URL != "" {
		addRow("URL", m.styles.CategoryApp.Copy().Foreground(m.styles.CategoryApp.GetForeground()).Background(lipgloss.NoColor{}).Render(d.URL))
	}

	detailsBox := m.styles.Panel.Width(m.width - 2).Render(detailsBuilder.String())

	// --- Footer ---
	footerBox := m.styles.Subtle.Render("Navigation: [Esc/q] Back to list")

	content := lipgloss.JoinVertical(lipgloss.Left, headerBox, detailsBox, footerBox)
	return lipgloss.Place(m.width, m.height, lipgloss.Left, lipgloss.Top, content)
}
