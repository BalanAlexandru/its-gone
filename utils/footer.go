package utils

import "github.com/BalanAlexandru/its-gone/styles"

func MakeFooter() string {
	return styles.InfoText.Render("Quit [q, ctrl+c] || Add task [a] || Select task [space]")
}
