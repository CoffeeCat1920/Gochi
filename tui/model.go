package tui

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/CoffeeCat1920/Gochi/entities"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor  int
	entries []entities.AppEntry
}

func newModel(store entities.AppStore) *model {
	return &model{
		cursor:  0,
		entries: store.Entries,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Gochi")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}
		case "enter", " ":
			cmdStr := strings.ReplaceAll(m.entries[m.cursor].Exec, "%u", "")
			parts := strings.Fields(cmdStr)
			cmd := exec.Command(parts[0], parts[1:]...)
			err := cmd.Start()
			if err != nil {
				fmt.Println("Error launching app:", err)
			}
			return m, tea.Quit

		}
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	for i, entry := range m.entries {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s.WriteString(fmt.Sprintf("%s %s\n", cursor, entry.Name))
	}

	return s.String()
}
