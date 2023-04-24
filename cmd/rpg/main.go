package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	lengthFlag := flag.Int("length", 12, "Set length of your password")
	uppercaseFlag := flag.Bool("uppercase", false, "Adds uppercase letters to your password (ABC)")
	lowercaseFlag := flag.Bool("lowercase", false, "Adds lowercase letters to your password (abc)")
	numberFlag := flag.Bool("number", false, "Adds numbers to your password (123)")
	specialCharFlag := flag.Bool("specialChar", false, "Adds special characters to your password ($%&)")

	var (
		letters      string
		numbers      string
		specialChars string
		charSet      string
	)

	// add uppercase to letters variable if uppercaseFlag's value is true
	if *uppercaseFlag {
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		*uppercaseFlag = true
		// add letters to charSet
		charSet += letters
	}
	// add lowercase to letters variable if lowercaseFlag's value is true
	if *lowercaseFlag {
		letters += "abcdefghijklmnopqrstuvwxyz"
		*lowercaseFlag = true
		// add letters to charSet
		charSet += letters
	}
	// add integers to numbers variable if numberFlag's value is true
	if *numberFlag {
		numbers += "0123456789"
		*numberFlag = true
		// add numbers to charSet
		charSet += numbers
	}
	// add special characters to specialCharFlag variable if specialCharFlag's value is true
	if *specialCharFlag {
		specialChars += "!@#$%^&*()-_=+,.?/:;{}[]"
		*specialCharFlag = true
		// add special characters to charSet
		charSet += specialChars
	}

	// creates an array of lengthFlag's length
	password := make([]byte, *lengthFlag)

	for i := 0; i < *lengthFlag; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// Print the password
	fmt.Println(string(password))
}
