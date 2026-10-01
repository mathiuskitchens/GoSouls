package ui

import (
	"charm.land/lipgloss/v2"
)

var HudBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7F1D1D")). // Dark blood-red border
	Foreground(lipgloss.Color("#F3F4F6")).       // Off-white text
	Padding(1, 2).                               // 1 line top/bottom, 2 spaces left/right
	Margin(1).                                   // 1 space margin around the box
	Border(lipgloss.ThickBorder())
