package models

import (
	"fmt"

	"github.com/BalanAlexandru/its-gone/styles"
	"github.com/BalanAlexandru/its-gone/utils"
)

func HomeView(model BubbleTeaModel) string {
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
	utils.HomeFooter()

	var options []utils.FooterOption = make([]utils.FooterOption, 0)

	if len(model.Items) > 0 {
		options = append(options, utils.DELETE_TASK_OPTION)
	}

	ui += styles.Footer.Render(utils.MakeFooter(options...))

	return styles.AppStyle.Render(ui)
}
