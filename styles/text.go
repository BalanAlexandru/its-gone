package styles

import gloss "github.com/charmbracelet/lipgloss"

var BoldText = gloss.NewStyle().
	Bold(true).
	Foreground(gloss.Color("#EBAF25"))

var InfoText = gloss.NewStyle().
	Foreground(gloss.Color("#B5AA99"))

var VerticalMargin = gloss.NewStyle().Margin(1, 0)
