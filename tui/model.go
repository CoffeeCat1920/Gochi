package tui

import (
	"fmt"
	"strings"

	"github.com/CoffeeCat1920/Gochi/loader"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	entries []loader.AppEntry
}

func newModel(entries []loader.AppEntry) *model {
	return &model{
		entries: entries,
	}
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
		case "ctrl+z":
			return m, tea.Suspend
		}
	}
	return m, nil
}

func (m model) View() string {
	for _, e := range m.entries {
		fmt.Printf("Name: %s, Command:%s \n", e.Name, e.Exec)
	}

	var b strings.Builder

	for _, e := range m.entries {
		b.WriteString(fmt.Sprintf("Name: %s, Command: %s\n", e.Name, e.Exec))
	}

	return b.String()
}
