package styles

import gloss "github.com/charmbracelet/lipgloss"

var AppStyle = gloss.NewStyle().Padding(2, 2).
	Background(gloss.Color("#2F4858")).
	Width(1920)

var Footer = gloss.NewStyle().MarginTop(1).
	Background(gloss.Color("#2F4858")).
	Border(gloss.NormalBorder(), true, false).
	BorderForeground(gloss.Color("#B5AA99"))

var Header = gloss.NewStyle().MarginBottom(1)
