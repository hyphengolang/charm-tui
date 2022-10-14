package page02

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hyphengolang/charm-tui/component/counter"
)

type Page struct {
	c counter.Counter
}

func New(n int) Page { return Page{c: counter.New(n)} }

func (pg Page) Init() tea.Cmd { return tea.Batch(pg.c.Init()) }

func (pg Page) View() string {
	var sb strings.Builder
	sb.WriteString("This is Page 2!!!")
	sb.WriteRune('\n')
	sb.WriteString(pg.c.View())
	return sb.String()
}

func (pg Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
