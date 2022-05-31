package styles

import gloss "github.com/charmbracelet/lipgloss"

var AppStyle = gloss.NewStyle().Padding(2, 2).
	Background(gloss.Color("#2F4858")).
	Width(256)

var Footer = gloss.NewStyle().MarginTop(1).
	Background(gloss.Color("#2F4858")).
	BorderStyle(gloss.NormalBorder()).
	BorderForeground(gloss.Color("#B5AA99")).
	BorderLeft(false).BorderRight(false).BorderTop(true).BorderBottom(true)

var Header = gloss.NewStyle().MarginBottom(1)
