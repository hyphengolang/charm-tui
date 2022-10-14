package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	page "github.com/hyphengolang/charm-tui/pages"
)

func main() {
	// new model
	p := tea.NewProgram(page.New())
	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}
}
