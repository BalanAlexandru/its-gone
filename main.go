package main

import (
	"log"

	"github.com/BalanAlexandru/its-gone/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(models.MakeDemoBubbleTeaModel())
	if err := program.Start(); err != nil {
		log.Fatal(err)
	}
}
