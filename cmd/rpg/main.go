package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// define command line flags
	lengthFlag := flag.Int("length", 8, "the length of the password")
	uppercaseFlag := flag.Bool("uppercase", false, "include uppercase letters in the password")
	lowercaseFlag := flag.Bool("lowercase", false, "include lowercase letters in the password")
	numberFlag := flag.Bool("number", false, "include numbers in the password")
	specialCharFlag := flag.Bool("special", false, "include special characters in the password")

	// parse command line flags
	flag.Parse()

	var charSet string

	// define character sets
	var letters string
	if *uppercaseFlag {
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		charSet += letters
	}
	if *lowercaseFlag {
		letters += "abcdefghijklmnopqrstuvwxyz"
		charSet += letters
	}
	var numbers string
	if *numberFlag {
		numbers = "0123456789"
		charSet += numbers
	}
	var specialChars string
	if *specialCharFlag {
		specialChars = "!@#$%^&*()-_=+,.?/:;{}[]"
		charSet += specialChars
	}

	// initialize the password variable
	password := make([]byte, *lengthFlag)

	// add characters to the password
	for i := 0; i < *lengthFlag; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// print the password
	fmt.Println(string(password))
}
