package ui

import (
	"charm.land/lipgloss/v2"
)

var HudBoxStyle = lipgloss.NewStyle().
	Bold(true).
	BorderForeground(lipgloss.Color("#7F1D1D")). // Dark blood-red border
	Foreground(lipgloss.Color("#F3F4F6")).       // Off-white text
	Padding(2).                                  // 1 line top/bottom, 2 spaces left/right
	Margin(2).                                   // 1 space margin around the box
	Border(myBorder).
	Width(50).
	Align(lipgloss.Center)

var myBorder = lipgloss.Border{
	Top:    ": ▬▬ι═══════ﺤ :",
	Bottom: ": ▬▬ι═══════ﺤ :",
	Left:   "|)|(",
	Right:  "|(|)",
}
