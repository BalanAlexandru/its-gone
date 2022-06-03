package styles

import gloss "github.com/charmbracelet/lipgloss"

var AppStyle = gloss.NewStyle().Padding(2, 2).
	Foreground(Text.GetForeground()).
	Background(gloss.Color(BACKGROUND)).
	Width(1920)

var Footer = gloss.NewStyle().MarginTop(1).
	Background(gloss.Color(BACKGROUND)).
	Border(gloss.NormalBorder(), true, false).
	BorderForeground(gloss.Color(MUTED))

var Header = gloss.NewStyle().MarginBottom(1)

var InputBackground = gloss.NewStyle().Background(gloss.Color(INPUT_BACKGROUND))
