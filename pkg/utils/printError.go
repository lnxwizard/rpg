package utils

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func PrintError(err error) {
	// Error message
	errorMessage := "Error: There was a problem opening repository in browser."

	// Marking word error message
	emphasized := marker.Mark(errorMessage, marker.MatchAll("Error:"), color.New(color.FgHiRed))

	// Printing error
	fmt.Println(emphasized)
}
