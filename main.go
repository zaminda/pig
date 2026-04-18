package main

import (
	"fmt"
	"os"

	"github.com/zaminda/pig/tui"

	tea "github.com/charmbracelet/bubbletea"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("pig " + version)
		return
	}

	p := tea.NewProgram(tui.NewModel(version))
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
