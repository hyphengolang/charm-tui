package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	page "github.com/hyphengolang/charm-tui/pages"
)

func main() {
	// new model
	app := page.New()

	p := tea.NewProgram(app)
	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}
}
