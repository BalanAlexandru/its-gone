package utils

import "github.com/BalanAlexandru/its-gone/styles"

var Footer string = ""

func MakeFooter() {
	Footer = styles.InfoText.Render("Quit [q, ctrl+c] || Add task [ctrl+a] || Select task [space]")
}

func AddTaskFooter() {
	Footer = styles.InfoText.Render("Quit [ctrl+c] || Go back [ctrl+b] || Save new task [enter]")
}

func UpdateFooter(message string) {
	Footer = styles.InfoText.Render(message)
}
