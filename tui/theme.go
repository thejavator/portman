package tui

import "github.com/charmbracelet/lipgloss"

// Theme defines the color palette for the UI
type Theme struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Text      lipgloss.Color
	Subtle    lipgloss.Color
	Error     lipgloss.Color
	
	CatApp    lipgloss.Color
	CatSys    lipgloss.Color
	CatDev    lipgloss.Color
	CatNet    lipgloss.Color
	CatOth    lipgloss.Color
	
	Background lipgloss.Color
}

// Styles holds all pre-compiled lipgloss styles based on a theme
type Styles struct {
	ActiveTab   lipgloss.Style
	InactiveTab lipgloss.Style
	Header      lipgloss.Style
	SelectedRow lipgloss.Style
	NormalRow   lipgloss.Style
	Subtle      lipgloss.Style
	
	CategoryApp lipgloss.Style
	CategorySys lipgloss.Style
	CategoryDev lipgloss.Style
	CategoryNet lipgloss.Style
	CategoryOth lipgloss.Style
	
	Title       lipgloss.Style
	Description lipgloss.Style
	Version     lipgloss.Style
	Error       lipgloss.Style
	
	Panel       lipgloss.Style
	ActivePanel lipgloss.Style
}

// Predefined themes
var Themes = map[string]Theme{
	"default": {
		Primary:    lipgloss.Color("46"),  // Green
		Secondary:  lipgloss.Color("240"), // Dark Gray
		Text:       lipgloss.Color("252"), // Off White
		Subtle:     lipgloss.Color("241"), // Medium Gray
		Error:      lipgloss.Color("196"), // Red
		CatApp:     lipgloss.Color("39"),  // Blue
		CatSys:     lipgloss.Color("208"), // Orange
		CatDev:     lipgloss.Color("42"),  // Green-Cyan
		CatNet:     lipgloss.Color("135"), // Purple
		CatOth:     lipgloss.Color("245"), // Gray
		Background: lipgloss.Color("236"), // Dark Background for badges
	},
	"dracula": {
		Primary:    lipgloss.Color("#FF79C6"), // Pink
		Secondary:  lipgloss.Color("#6272A4"), // Purple Gray
		Text:       lipgloss.Color("#F8F8F2"), // White
		Subtle:     lipgloss.Color("#44475A"), // Dark Gray
		Error:      lipgloss.Color("#FF5555"), // Red
		CatApp:     lipgloss.Color("#8BE9FD"), // Cyan
		CatSys:     lipgloss.Color("#FFB86C"), // Orange
		CatDev:     lipgloss.Color("#50FA7B"), // Green
		CatNet:     lipgloss.Color("#BD93F9"), // Purple
		CatOth:     lipgloss.Color("#F1FA8C"), // Yellow
		Background: lipgloss.Color("#282A36"), // Dark Background
	},
	"nord": {
		Primary:    lipgloss.Color("#88C0D0"), // Frost Blue
		Secondary:  lipgloss.Color("#4C566A"), // Nord Gray
		Text:       lipgloss.Color("#ECEFF4"), // Snow White
		Subtle:     lipgloss.Color("#3B4252"), // Dark Gray
		Error:      lipgloss.Color("#BF616A"), // Red
		CatApp:     lipgloss.Color("#5E81AC"), // Deep Blue
		CatSys:     lipgloss.Color("#D08770"), // Orange
		CatDev:     lipgloss.Color("#A3BE8C"), // Green
		CatNet:     lipgloss.Color("#B48EAD"), // Purple
		CatOth:     lipgloss.Color("#EBCB8B"), // Yellow
		Background: lipgloss.Color("#2E3440"), // Dark Background
	},
}

// GetStyles returns a fully instantiated Styles struct based on the requested theme
func GetStyles(themeName string) Styles {
	theme, exists := Themes[themeName]
	if !exists {
		theme = Themes["default"]
	}

	return Styles{
		ActiveTab:   lipgloss.NewStyle().Foreground(theme.Background).Background(theme.Primary).Bold(true).Padding(0, 1),
		InactiveTab: lipgloss.NewStyle().Foreground(theme.Secondary).Padding(0, 1),
		Header:      lipgloss.NewStyle().Foreground(theme.Text).Bold(true),
		SelectedRow: lipgloss.NewStyle().Foreground(theme.Primary).Bold(true),
		NormalRow:   lipgloss.NewStyle().Foreground(theme.Text),
		Subtle:      lipgloss.NewStyle().Foreground(theme.Subtle),
		
		CategoryApp: lipgloss.NewStyle().Foreground(theme.CatApp).Background(theme.Background).Padding(0, 1),
		CategorySys: lipgloss.NewStyle().Foreground(theme.CatSys).Background(theme.Background).Padding(0, 1),
		CategoryDev: lipgloss.NewStyle().Foreground(theme.CatDev).Background(theme.Background).Padding(0, 1),
		CategoryNet: lipgloss.NewStyle().Foreground(theme.CatNet).Background(theme.Background).Padding(0, 1),
		CategoryOth: lipgloss.NewStyle().Foreground(theme.CatOth).Background(theme.Background).Padding(0, 1),
		
		Title:       lipgloss.NewStyle().Foreground(theme.Primary).Bold(true),
		Description: lipgloss.NewStyle().Foreground(theme.Text),
		Version:     lipgloss.NewStyle().Foreground(theme.Secondary),
		Error:       lipgloss.NewStyle().Foreground(theme.Error).Bold(true),
		
		Panel:       lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(theme.Subtle).Padding(0, 1),
		ActivePanel: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(theme.Primary).Padding(0, 1),
	}
}
