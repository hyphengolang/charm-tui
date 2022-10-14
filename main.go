package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// new model
	app := NewDocument()

	p := tea.NewProgram(app)
	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}
}
