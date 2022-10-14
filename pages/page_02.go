package page

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hyphengolang/charm-tui/component/counter"
)

type PageTwo struct {
	c counter.Counter
}

func pageTwo(n int) PageTwo { return PageTwo{c: counter.New(n)} }

func (pg PageTwo) Init() tea.Cmd { return tea.Batch(pg.c.Init()) }

func (pg PageTwo) View() string {
	var sb strings.Builder
	sb.WriteString("Page 2 is in progress")
	sb.WriteRune('\n')
	sb.WriteString(pg.c.View())
	return sb.String()
}

func (pg PageTwo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// case tea.KeyBackspace:
		// 	return pg, levelUp()
		}
	}

	var cmd tea.Cmd
	pg.c, cmd = pg.c.Update(msg)

	return pg, tea.Batch(cmd)
}
