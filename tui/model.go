package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var items = []string{
	"Apples",
	"Bananas",
	"Cherries",
	"Dates",
	"Elderberries",
}

type Model struct {
	Items    []string
	Cursor   int
	Quitting bool
}

func NewModel() Model {
	return Model{Items: items}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.Quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString("🐷 Pig — Pick a fruit!\n\n")

	for i, item := range m.Items {
		cursor := "  "
		if m.Cursor == i {
			cursor = "▸ "
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, item)
	}

	b.WriteString("\n  j/k or ↑/↓ to move • q to quit\n")
	return b.String()
}
