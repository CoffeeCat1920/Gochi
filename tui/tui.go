package tui

import (
	"log"

	"github.com/CoffeeCat1920/Gochi/loader"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(entries []loader.AppEntry) {
	m := newModel(entries)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
