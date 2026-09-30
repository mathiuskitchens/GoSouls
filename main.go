package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	"os"
)

// lip gloss styling section here...
var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("FAFAFA")).
	Background(lipgloss.Color("7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

var hudBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7F1D1D")). // Dark blood-red border
	Foreground(lipgloss.Color("#F3F4F6")).       // Off-white text
	Padding(1, 2).                               // 1 line top/bottom, 2 spaces left/right
	Margin(1).                                   // 1 space margin around the box
	Border(lipgloss.ThickBorder())

// This is a struct that outlines a blueprint for what a character is
type Character struct {
	name            string
	currentHealth   int
	maxHealth       int
	equippedWeapons []string
}

// root state container...we nest Character in since we outlined it above
type model struct {
	player    Character
	lastEvent string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "d":
			m.player.currentHealth -= 5
			h := m.player.currentHealth
			switch {
			case h <= 0:
				m.lastEvent = "YOU DIED. \n"
				return m, tea.Quit
			case h > 0:
				m.lastEvent = "You took 5 damage! \n"
			}
		case "b":
			m.player.currentHealth -= 10
			h := m.player.currentHealth
			switch {
			case h <= 0:
				m.lastEvent = "YOU DIED! REALLY! \n"
				return m, tea.Quit
			case h > 0:
				m.lastEvent = "You took 10 damage! \n"

			}
		}

	}
	return m, nil
}

func (m model) View() tea.View {
	content := fmt.Sprintf(
		"%sYour health is %d/%d. \nPress 'q' to quit.\n",
		m.lastEvent,
		m.player.currentHealth,
		m.player.maxHealth,
	)

	styledContent := hudBoxStyle.Render(content)
	return tea.NewView(styledContent)
}

func main() {
	p := tea.NewProgram(model{
		player: Character{
			name:          "Mathius",
			currentHealth: 20,
			maxHealth:     20,
		},
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
