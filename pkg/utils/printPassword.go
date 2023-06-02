package utils

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func PrintPassword(password string) {
	// Password message
	passwordMessage := "Your RP: " + password

	// Marking word "Your RP:"
	emphasized := marker.Mark(passwordMessage, marker.MatchAll("Your RP:"), color.New(color.FgGreen))

	// Printing password
	fmt.Println(emphasized)
}
