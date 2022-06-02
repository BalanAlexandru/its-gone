package utils

import (
	"time"

	"github.com/BalanAlexandru/its-gone/styles"
)

const name = "Alexandru"

func MakeHeader() string {
	return styles.Header.Render("Good " + FindDayTime() + ", " + name + "!")
}

func UpdateHeader(message string) string {
	return styles.Header.Render(message)
}

func FindDayTime() string {
	now := time.Now()

	if now.Hour() < 12 {
		return "morning"
	} else if now.Hour() < 17 {
		return "afternoon"
	} else if now.Hour() < 20 {
		return "evening"
	} else {
		return "night"
	}

}
