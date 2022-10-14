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
	var cmd tea.Cmd
	pg.t, cmd = pg.t.Update(msg)

	return pg, tea.Batch(cmd)
}

func (pg Page) View() string {
	var sb strings.Builder
	sb.WriteString("This is Page 1!!!")
	sb.WriteString(pg.t.View())
	return sb.String()
}
