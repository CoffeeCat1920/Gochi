package tui

import (
	"log"

	"github.com/CoffeeCat1920/Gochi/entities"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(store entities.AppStore) {
	m := newModel(store)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
