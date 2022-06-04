package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ViewType int64

const (
	Home ViewType = iota
	AddTask
)

type BubbleTeaModel struct {
	tea.Model
	TextInput   TextInputModel
	CurrentView ViewType
	Items       []Task
	Cursor      int
	Selected    map[int]struct{}
}

var initialModel BubbleTeaModel = BubbleTeaModel{
	TextInput:   MakeTextInputModel(),
	CurrentView: Home,
	Items:       make([]Task, 0),
	Selected:    make(map[int]struct{}),
}

func MakeDemoBubbleTeaModel() BubbleTeaModel {
	return initialModel
}

func (model BubbleTeaModel) Init() tea.Cmd {
	return tea.Batch(
		model.TextInput.Init(),
	)
}

func (model BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var commands []tea.Cmd = make([]tea.Cmd, 5)

	switch model.CurrentView {
	case Home:
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

			case " ":
				_, ok := model.Selected[model.Cursor]
				if ok {
					delete(model.Selected, model.Cursor)
				} else {
					model.Selected[model.Cursor] = struct{}{}
				}

			case tea.KeyCtrlA.String():
				model.CurrentView = AddTask

			case "d":
				if len(model.Selected) != 0 {
					for k := range model.Selected {
						if model.Items[k].State == Done {
							model.Items[k].State = New
						} else {
							model.Items[k].State = Done
						}
					}
				}
			}
		}

	case AddTask:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case tea.KeyCtrlC.String():
				commands = append(commands, tea.Quit)

			case tea.KeyEnter.String():
				if strings.TrimSpace(model.TextInput.Value()) != "" {
					newTask := Task{
						State: New,
						Name:  model.TextInput.Value(),
					}
					model.Items = append(model.Items, newTask)
					model.TextInput = MakeTextInputModel()
				}

				model.CurrentView = Home
				return model, tea.Batch(commands...)

			case tea.KeyCtrlB.String():
				model.CurrentView = Home
			}
		}

		model.TextInput, cmd = model.TextInput.Update(msg)
		commands = append(commands, cmd)
	}

	return model, tea.Batch(commands...)
}

func (model BubbleTeaModel) View() string {
	switch model.CurrentView {
	case Home:
		return HomeView(model)
	case AddTask:
		return AddTaskView(model)
	}

	return HomeView(model)
}
