package page01

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hyphengolang/charm-tui/component/tabs"
)

type Page struct {
	t tabs.Tabs
}

func New() Page { return Page{t: tabs.New()} }

func (pg Page) Init() tea.Cmd { return tea.Batch(pg.t.Init()) }

func (pg Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			return pg, nil
		}
	}

	pg.t, cmd = pg.t.Update(msg)
	cmds = append(cmds, cmd)

	return pg, tea.Batch(cmds...)
}

func (pg Page) View() string {
	var sb strings.Builder
	sb.WriteString("This is Page 1!!!")
	sb.WriteString(pg.t.View())
	return sb.String()
}
