package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Value int
}

type App struct {
	pg Page

	mpg map[Page]tea.Model
}

func NewDocument() App {
	m := App{
		mpg: map[Page]tea.Model{
			PgOne: NewPageOne(5),
			PgTwo: NewPageTwo(),
		},
	}
	return m
}

func (a App) Init() tea.Cmd { return tea.Batch(tea.EnterAltScreen) }

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return a, tea.Quit
		case tea.KeyTab:
			a.pg.Next()
		}
	}

	if m, ok := a.mpg[a.pg]; ok {
		a.mpg[a.pg], cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return a, tea.Batch(cmds...)
}

func (a App) focusedPage() tea.Model { return a.mpg[a.pg] }

func (a App) View() string {
	var sb strings.Builder
	sb.WriteString("This is the top-level application \u2022\n\n")
	sb.WriteString(a.focusedPage().View())
	return sb.String()
}
