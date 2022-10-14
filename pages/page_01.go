package page

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hyphengolang/charm-tui/component/tabs"
)

type PageOne struct {
	t tabs.Tabs
}

func pageOne() PageOne { return PageOne{t: tabs.New()} }

func (pg PageOne) Init() tea.Cmd { return tea.Batch(pg.t.Init()) }

func (pg PageOne) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// case tea.KeyBackspace:
		// 	return pg, levelUp()
		}
	}

	pg.t, cmd = pg.t.Update(msg)
	cmds = append(cmds, cmd)

	return pg, tea.Batch(cmds...)
}

func (pg PageOne) View() string {
	var sb strings.Builder
	sb.WriteString("This is Page 1!!!")
	sb.WriteString(pg.t.View())
	return sb.String()
}
