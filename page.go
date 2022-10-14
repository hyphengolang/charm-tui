package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Page uint

const (
	PgOne Page = iota
	PgTwo
)

func (pg *Page) Next() {
	switch *pg {
	case PgTwo:
		*pg = 0
	default:
		*pg++
	}
}

func (pg *Page) Prev() {
	switch *pg {
	case 0:
		*pg = PgTwo
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

var _ tea.Model = (*PageOne)(nil)
var _ tea.Model = (*PageTwo)(nil)

type PageOne struct {
	c Counter
}

func NewPageOne(n int) PageOne { return PageOne{c: NewCounter(n)} }

func (pg PageOne) Init() tea.Cmd { return tea.Batch() }

func (pg PageOne) View() string {
	// header
	// main
	// footer
	var sb strings.Builder
	sb.WriteString("This is Foo!!!\n\n")
	sb.WriteString(pg.c.View())
	return sb.String()
}

func (pg PageOne) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		//...
		}
	}

	var cmd tea.Cmd
	pg.c, cmd = pg.c.Update(msg)

	return pg, tea.Batch(cmd)
}

type PageTwo struct{}

func NewPageTwo() PageTwo { return PageTwo{} }

func (pg PageTwo) Init() tea.Cmd { return tea.Batch() }

func (pg PageTwo) View() string {
	var sb strings.Builder
	sb.WriteString("This is Bar!!!\n\n")
	return sb.String()
}

func (pg PageTwo) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return pg, nil }
