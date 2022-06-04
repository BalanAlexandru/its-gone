package styles

import gloss "github.com/charmbracelet/lipgloss"

var Text = gloss.NewStyle().Foreground(gloss.Color(FOREGROUND))

var BoldText = Text.Copy().
	Bold(true)

var DoneText = BoldText.Copy().Foreground(gloss.Color(SUCCESS)).
	Background(gloss.Color(BACKGROUND)).
	Strikethrough(true)

var InfoText = Text.Copy().
	Foreground(gloss.Color(INFO))

var MutedText = Text.Copy().
	Foreground(gloss.Color(MUTED))

var VerticalMargin = gloss.NewStyle().Margin(1, 0)
