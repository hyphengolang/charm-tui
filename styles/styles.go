package styles

import (
	css "github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	Color   = termenv.EnvColorProfile().Color
	Keyword = termenv.Style{}.Foreground(Color("204")).Background(Color("235")).Styled
	Help    = termenv.Style{}.Foreground(Color("241")).Styled
	H1      = css.NewStyle().Foreground(css.Color("204")).Background(css.Color("235"))
)
