package styles

import "github.com/muesli/termenv"

var (
	Color   = termenv.EnvColorProfile().Color
	Keyword = termenv.Style{}.Foreground(Color("204")).Background(Color("235")).Styled
	Help    = termenv.Style{}.Foreground(Color("241")).Styled
)
