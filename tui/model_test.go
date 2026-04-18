package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	m := NewModel("test")
	if len(m.Items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(m.Items))
	}
	if m.Cursor != 0 {
		t.Fatalf("expected cursor at 0, got %d", m.Cursor)
	}
}

func TestInit(t *testing.T) {
	m := NewModel("test")
	if cmd := m.Init(); cmd == nil {
		t.Fatal("Init should return a command for update check")
	}
}

func TestUpdateAvailableMsg(t *testing.T) {
	m := NewModel("test")
	updated, cmd := m.Update(updateAvailableMsg("update available: v0.2.0"))
	um := updated.(Model)
	if um.UpdateNotice != "update available: v0.2.0" {
		t.Fatalf("expected update notice, got %q", um.UpdateNotice)
	}
	if cmd != nil {
		t.Fatal("expected no command after update notice")
	}
}

func TestViewShowsUpdateNotice(t *testing.T) {
	m := NewModel("0.1.0")
	m.UpdateNotice = "update available: v0.2.0"
	v := m.View()
	if !strings.Contains(v, "update available: v0.2.0") {
		t.Fatal("expected update notice in view")
	}
}

func TestNormalise(t *testing.T) {
	if normalise("v1.2.3") != "1.2.3" {
		t.Fatal("expected v prefix stripped")
	}
	if normalise("1.2.3") != "1.2.3" {
		t.Fatal("expected no change without v prefix")
	}
}

func TestCursorDown(t *testing.T) {
	m := NewModel("test")
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	um := updated.(Model)
	if um.Cursor != 1 {
		t.Fatalf("expected cursor at 1, got %d", um.Cursor)
	}
}

func TestCursorUp(t *testing.T) {
	m := NewModel("test")
	m.Cursor = 2
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	um := updated.(Model)
	if um.Cursor != 1 {
		t.Fatalf("expected cursor at 1, got %d", um.Cursor)
	}
}

func TestCursorDoesNotGoBelowZero(t *testing.T) {
	m := NewModel("test")
	m.Cursor = 0
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	um := updated.(Model)
	if um.Cursor != 0 {
		t.Fatalf("expected cursor at 0, got %d", um.Cursor)
	}
}

func TestCursorDoesNotExceedItems(t *testing.T) {
	m := NewModel("test")
	m.Cursor = len(m.Items) - 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	um := updated.(Model)
	if um.Cursor != len(m.Items)-1 {
		t.Fatalf("expected cursor at %d, got %d", len(m.Items)-1, um.Cursor)
	}
}

func TestQuit(t *testing.T) {
	m := NewModel("test")
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	um := updated.(Model)
	if !um.Quitting {
		t.Fatal("expected quitting to be true")
	}
	if cmd == nil {
		t.Fatal("expected quit command")
	}
}

func TestViewShowsCursor(t *testing.T) {
	m := NewModel("test")
	m.Cursor = 2
	v := m.View()
	lines := strings.Split(v, "\n")

	// Find the line with Cherries (cursor=2, 3rd item)
	found := false
	for _, line := range lines {
		if strings.Contains(line, "▸") && strings.Contains(line, "Cherries") {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected cursor on Cherries")
	}
}

func TestViewEmptyOnQuit(t *testing.T) {
	m := NewModel("test")
	m.Quitting = true
	if v := m.View(); v != "" {
		t.Fatalf("expected empty view on quit, got %q", v)
	}
}
