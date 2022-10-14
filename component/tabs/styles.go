package tabs

import css "github.com/charmbracelet/lipgloss"

func tabBorderWithBottom(left, middle, right string) css.Border {
	border := css.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = css.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = css.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = css.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = css.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(css.Center).Border(css.NormalBorder()).UnsetBorderTop()
)
