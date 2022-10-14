package page

import (
	tea "github.com/charmbracelet/bubbletea"
	styles "github.com/charmbracelet/lipgloss"
)

type levelMsg struct {
}

func levelUp() tea.Cmd {
	return func() tea.Msg { return nil }
}

func levelDown() tea.Cmd {
	return func() tea.Msg { return nil }
}

type State struct {
	Value int
}

type Document struct {
	c, w, h int

	pg map[int]tea.Model
}

func New() Document {
	m := Document{
		pg: map[int]tea.Model{
			0: pageOne(),
			1: pageTwo(5),
		},
	}
	return m
}

func (doc Document) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.EnterAltScreen)

	for _, m := range doc.pg {
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
		doc.h, doc.w = msg.Height, msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return doc, tea.Quit
		case tea.KeyTab:
			/* return */ doc.Next()
		case tea.KeyShiftTab:
			/* return */ doc.Prev()
			// case tea.KeyBackspace:
			// 	return doc, nil
		}
	case levelMsg:
		return doc, nil
	}

	if m, ok := doc.pg[doc.c]; ok {
		doc.pg[doc.c], cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return doc, tea.Batch(cmds...)
}

func (doc Document) View() string {
	styles := styles.NewStyle().BorderStyle(styles.RoundedBorder())

	return styles.Render(doc.Current().View())
}

func (doc Document) Current() tea.Model { return doc.pg[doc.c] }

func (doc *Document) Next() (tea.Model, tea.Cmd) {
	if max := len(doc.pg) - 1; doc.c == max {
		doc.c = 0
	} else {
		doc.c++
	}

	return doc.pg[doc.c], levelDown()
}

func (doc *Document) Prev() (tea.Model, tea.Cmd) {
	if max := len(doc.pg) - 1; doc.c == 0 {
		doc.c = max
	} else {
		doc.c--
	}

	return doc.pg[doc.c], levelDown()
}
