package page01

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Page struct{}

func New() Page { return Page{} }

func (pg Page) Init() tea.Cmd { return tea.Batch() }

func (pg Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return pg, nil }

func (pg Page) View() string {
	var sb strings.Builder
	sb.WriteString("This is Page 1!!!\n\n")
	return sb.String()
}
