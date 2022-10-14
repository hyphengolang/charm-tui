package page

import (
	tea "github.com/charmbracelet/bubbletea"
	page01 "github.com/hyphengolang/charm-tui/pages/page-01"
	page02 "github.com/hyphengolang/charm-tui/pages/page-02"
)

type State struct {
	Value int
}

type Document struct {
	c uint

	pgs map[uint]tea.Model
}

func New() Document {
	m := Document{
		pgs: map[uint]tea.Model{
			0: page01.New(),
			1: page02.New(5),
		},
	}
	return m
}

func (doc Document) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.EnterAltScreen)

	for _, m := range doc.pgs {
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
			doc.Next()
		case tea.KeyShiftTab:
			doc.Prev()
		}
	}

	if m, ok := doc.pgs[doc.c]; ok {
		doc.pgs[doc.c], cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return doc, tea.Batch(cmds...)
}

func (doc Document) View() string { return doc.Current().View() }

func (doc Document) Current() tea.Model { return doc.pgs[doc.c] }

func (doc *Document) Next() {
	if max := uint(len(doc.pgs) - 1); doc.c == max {
		doc.c = 0
	} else {
		doc.c++
	}
}

func (doc *Document) Prev() {
	if max := uint(len(doc.pgs) - 1); doc.c == 0 {
		doc.c = max
	} else {
		doc.c--
	}
}

// TODO enum could be removed
// as it is nothing more than a counter
type Page uint

const (
	One Page = iota
	Two
)

func (pg *Page) Next() {
	switch *pg {
	case Two:
		*pg = 0
	default:
		*pg++
	}
}

func (pg *Page) Prev() {
	switch *pg {
	case 0:
		*pg = Two
	default:
		*pg--
	}
}

func (pg Page) String() string {
	return [...]string{
		"page_one",
		"page_two",
	}[pg]
}
