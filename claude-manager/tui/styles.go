package tui

import "github.com/charmbracelet/lipgloss"

var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	dimmed    = lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(highlight).
			Padding(0, 1)

	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(highlight).
			Padding(0, 2)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(dimmed).
				Padding(0, 2)

	selectedStyle = lipgloss.NewStyle().
			Foreground(special).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"})

	detailStyle = lipgloss.NewStyle().
			Foreground(dimmed).
			PaddingLeft(4)

	helpStyle = lipgloss.NewStyle().
			Foreground(dimmed).
			PaddingTop(1)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"})

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555"))

	inputStyle = lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true)
)
