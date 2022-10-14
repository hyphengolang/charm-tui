package styles

import (
	css "github.com/charmbracelet/lipgloss"
)

var (
	Keyword = css.NewStyle().Foreground(css.Color("204")).Background(css.Color("235")).Render
	Help    = css.NewStyle().Foreground(css.Color("241")).Render
	H1      = css.NewStyle().Foreground(css.Color("204")).Background(css.Color("235")).Render
)
