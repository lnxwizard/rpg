package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/AlperAkca79/rpg/pkg/utils"
	"github.com/cli/browser"
)

func main() {
	// Define command line flags (length, uppercase letter, lowercase letter, number and special characters)
	lengthFlag := flag.Int("length", 8, "The length of the password")
	uppercaseFlag := flag.Bool("uppercase", false, "Include uppercase letters in the password")
	lowercaseFlag := flag.Bool("lowercase", false, "Include lowercase letters in the password")
	numberFlag := flag.Bool("number", false, "Include numbers in the password")
	specialCharFlag := flag.Bool("special", false, "Include special characters in the password")
	repoFlag := flag.Bool("repo", false, "Opens GitHub repository of rpg in your browser.")

	// Parse command line flags
	flag.Parse()

	// Define uppercase letters, lowercase letters, numbers and special characters
	var (
		charSet      string
		letters      string
		numbers      string
		specialChars string
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: rpg [options]")
	}

	switch {
	case *uppercaseFlag:
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		charSet += letters
	case *lowercaseFlag:
		letters += "abcdefghijklmnopqrstuvwxyz"
		charSet += letters
	case *numberFlag:
		numbers = "0123456789"
		charSet += numbers
	case *specialCharFlag:
		specialChars = "!@#$%^&*()-_=+,.?/:;{}[]<>|"
		charSet += specialChars
	case *repoFlag:
		err := browser.OpenURL("https://github.com/AlperAkca79/rpg")
		if err != nil {
			utils.PrintError(err)
		}
	}

	// Initialize the password variable
	password := make([]byte, *lengthFlag)

	// Add characters to the password
	for i := 0; i < *lengthFlag; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// Printing Password
	utils.PrintPassword(string(password))
}
