package models

import "github.com/BalanAlexandru/its-gone/styles"

type State int64

const (
	New State = iota
	Active
	Done
)

type Task struct {
	Name        string
	Description string
	State       State
}

func (task Task) GlossName() string {
	if task.State == Done {
		return styles.DoneText.Render(task.Name)
	}

	return styles.BoldText.Render(task.Name)
}

func (task *Task) ToggleCompletion() {
	if task.State == Done {
		task.State = New
	} else {
		task.State = Done
	}
}
