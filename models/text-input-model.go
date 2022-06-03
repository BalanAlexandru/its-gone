package models

import (
	"github.com/BalanAlexandru/its-gone/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputModel struct {
	tea.Model
	TextInput textinput.Model
}

func MakeTextInputModel() TextInputModel {
	model := TextInputModel{}
	model.TextInput = textinput.New()
	model.TextInput.Placeholder = "Type the name of the task..."
	model.TextInput.PlaceholderStyle = styles.MutedText
	model.TextInput.Focus()

	return model
}

func (model TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (model TextInputModel) Update(msg tea.Msg) (TextInputModel, tea.Cmd) {
	var cmd tea.Cmd

	model.TextInput, cmd = model.TextInput.Update(msg)
	return model, cmd
}

func (model TextInputModel) View() string {
	return styles.VerticalMargin.Render(styles.InputBackground.Render(model.TextInput.View()))
}

func (model TextInputModel) Value() string {
	return model.TextInput.Value()
}

func (model TextInputModel) Clear() {
	model.TextInput.Reset()
}
