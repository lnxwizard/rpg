package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// define command line flags (length, uppercase letter, lowercase letter, number and special characters)
	lengthFlag := flag.Int("length", 8, "the length of the password")
	uppercaseFlag := flag.Bool("uppercase", false, "include uppercase letters in the password")
	lowercaseFlag := flag.Bool("lowercase", false, "include lowercase letters in the password")
	numberFlag := flag.Bool("number", false, "include numbers in the password")
	specialCharFlag := flag.Bool("special", false, "include special characters in the password")

	// parse command line flags
	flag.Parse()

	// define uppercase letters, lowercase letters, numbers and special characters
	var (
		charSet      string
		letters      string
		numbers      string
		specialChars string
	)

	// if uppercaseFlag is true, put uppercase letters inside the password
	if *uppercaseFlag {
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		charSet += letters
	}

	// if lowercaseFlag is true, put lowercase letters inside the password
	if *lowercaseFlag {
		letters += "abcdefghijklmnopqrstuvwxyz"
		charSet += letters
	}

	// if numberFlag is true, put number inside the password
	if *numberFlag {
		numbers = "0123456789"
		charSet += numbers
	}

	// if specialCharFlag is true, put special characters inside the password
	if *specialCharFlag {
		specialChars = "!@#$%^&*()-_=+,.?/:;{}[]<>|"
		charSet += specialChars
	}

	// initialize the password variable
	password := make([]byte, *lengthFlag)

	// add characters to the password
	for i := 0; i < *lengthFlag; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// print the password
	fmt.Println("Your RP:", string(password))
}
