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
	return styles.BoldText.Render(task.Name)
}
