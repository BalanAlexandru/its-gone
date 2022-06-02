package models

import (
	"github.com/BalanAlexandru/its-gone/styles"
	"github.com/BalanAlexandru/its-gone/utils"
)

func AddTaskView(model BubbleTeaModel) string {
	// Header
	ui := utils.UpdateHeader("Adding a new task")

	ui += model.TextInput.View()

	// Footer
	utils.AddTaskFooter()
	ui += styles.Footer.Render(utils.Footer)

	return styles.AppStyle.Render(ui)
}
