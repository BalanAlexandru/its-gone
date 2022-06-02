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
	utils.MakeFooter()
	ui += styles.Footer.Render(utils.Footer)

	return styles.AppStyle.Render(ui)
}
