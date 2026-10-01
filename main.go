package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"github.com/mathiuskitchens/gosouls/internal/game"
	"github.com/mathiuskitchens/gosouls/internal/ui"
	"os"
)

// root state container...we nest Character in since we outlined it above
type model struct {
	player    game.Character
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
			m.player.CurrentHealth -= 5
			h := m.player.CurrentHealth
			switch {
			case h <= 0:
				m.lastEvent = "YOU DIED. \n"
				return m, tea.Quit
			case h > 0:
				m.lastEvent = "You took 5 damage! \n"
			}
		case "b":
			m.player.CurrentHealth -= 10
			h := m.player.CurrentHealth
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
		m.player.CurrentHealth,
		m.player.MaxHealth,
	)

	styledContent := ui.HudBoxStyle.Render(content)
	return tea.NewView(styledContent)
}

func main() {
	p := tea.NewProgram(model{
		player: game.Character{
			Name:          "Mathius",
			CurrentHealth: 20,
			MaxHealth:     20,
		},
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
