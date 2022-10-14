package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// var _ tea.Model = (*Counter)(nil)

type Counter struct{ Value int }

func NewCounter(n int) Counter { return Counter{n} }

func (c Counter) Init() tea.Cmd { return tea.Batch() }

func (c Counter) View() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Count: %d\n", c.Value))
	return sb.String()
}

func (c Counter) Update(msg tea.Msg) (Counter, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			c.Value += 1
		case tea.KeyDown:
			c.Value -= 1
		}
	}

	return c, tea.Batch()
}
