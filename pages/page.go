package page

import (
	tea "github.com/charmbracelet/bubbletea"
	css "github.com/charmbracelet/lipgloss"
	page01 "github.com/hyphengolang/charm-tui/pages/page-01"
	page02 "github.com/hyphengolang/charm-tui/pages/page-02"
)

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
			0: page01.New(),
			1: page02.New(5),
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
			doc.Next()
		case tea.KeyShiftTab:
			doc.Prev()
		}
	}

	if m, ok := doc.pg[doc.c]; ok {
		doc.pg[doc.c], cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return doc, tea.Batch(cmds...)
}

func (doc Document) View() string {
	styles := css.NewStyle().BorderStyle(css.RoundedBorder())

	return styles.Render(doc.Current().View())
}

func (doc Document) Current() tea.Model { return doc.pg[doc.c] }

func (doc *Document) Next() {
	if max := len(doc.pg) - 1; doc.c == max {
		doc.c = 0
	} else {
		doc.c++
	}
}

func (doc *Document) Prev() {
	if max := len(doc.pg) - 1; doc.c == 0 {
		doc.c = max
	} else {
		doc.c--
	}
}
