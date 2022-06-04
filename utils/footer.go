package utils

import "github.com/BalanAlexandru/its-gone/styles"

type FooterOption string

const MARK_DONE_OPTION FooterOption = "Toggle done/undone [d]"

const home_footer string = "Quit [q, ctrl+c] || Add task [ctrl+a] || Select task [space]"
const add_task_footer string = "Quit [ctrl+c] || Go back [ctrl+b] || Save new task [enter]"

var footer string = home_footer

func MakeFooter(options ...FooterOption) string {
	for _, option := range options {
		footer += " || " + string(option)
	}

	return styles.InfoText.Render(footer)
}

func HomeFooter() {
	footer = home_footer
}

func AddTaskFooter() {
	footer = add_task_footer
}

func UpdateFooter(message string) {
	footer = styles.InfoText.Render(message)
}
