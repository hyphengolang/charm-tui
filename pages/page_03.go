package page

import tea "github.com/charmbracelet/bubbletea"

type PageThree struct{}

func pageThree() PageThree {
	return PageThree{}
}

func (pg PageThree) Init() tea.Cmd { return tea.Batch() }

func (pg PageThree) View() string { return "Page 3 is in progress" }

func (pg PageThree) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return pg, tea.Batch()
}
