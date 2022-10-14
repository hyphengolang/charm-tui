package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Value int
}

type Document struct {
	pg Page

	body map[Page]tea.Model
}

func NewDocument() Document {
	m := Document{
		body: map[Page]tea.Model{
			PgOne: NewPageOne(5),
			PgTwo: NewPageTwo(),
		},
	}
	return m
}

func (doc Document) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.EnterAltScreen)

	for _, m := range doc.body {
		cmds = append(cmds, m.Init())
	}

	return tea.Batch(cmds...)
}

func (doc Document) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		_, _ = msg.Height, msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return doc, tea.Quit
		case tea.KeyTab:
			doc.pg.Next()
		case tea.KeyShiftTab:
			doc.pg.Prev()
		}
	}

	if m, ok := doc.body[doc.pg]; ok {
		doc.body[doc.pg], cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return doc, tea.Batch(cmds...)
}

func (doc Document) View() string { return doc.body[doc.pg].View() }
