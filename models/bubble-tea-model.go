package models

import (
	"fmt"

	"github.com/BalanAlexandru/its-gone/styles"
	"github.com/BalanAlexandru/its-gone/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleTeaModel struct {
	tea.Model
	Items    []Task
	Cursor   int
	Selected map[int]struct{}
}

var initialModel BubbleTeaModel = BubbleTeaModel{
	Items:    make([]Task, 0),
	Selected: make(map[int]struct{}),
}

func MakeDemoBubbleTeaModel() BubbleTeaModel {
	initialModel.Items = append(initialModel.Items, Task{Name: "Buy carrots", State: New})
	initialModel.Items = append(initialModel.Items, Task{Name: "Wash car", State: New})

	return initialModel
}

func (model BubbleTeaModel) Init() tea.Cmd {
	return nil
}

func (model BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return model, tea.Quit

		case "up", "k":
			if model.Cursor > 0 {
				model.Cursor--
			}

		case "down", "j":
			if model.Cursor < len(model.Items)-1 {
				model.Cursor++
			}

		case "enter", " ":
			_, ok := model.Selected[model.Cursor]
			if ok {
				delete(model.Selected, model.Cursor)
			} else {
				model.Selected[model.Cursor] = struct{}{}
			}
		}
	}

	return model, nil
}

func (model BubbleTeaModel) View() string {
	// Header
	ui := utils.MakeHeader()

	for i, choice := range model.Items {
		cursor := " "
		if model.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := model.Selected[i]; ok {
			checked = "x"
		}

		ui += styles.VerticalMargin.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice.GlossName()))
	}

	// Footer
	ui += styles.Footer.Render(utils.MakeFooter())

	return styles.AppStyle.Render(ui)
}
