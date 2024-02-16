package main

import "unicode"

//Have the function CodelandUsernameValidation(str) take the str parameter being passed and determine if the string is a valid username according to the following rules:
//
//1. The username is between 4 and 25 characters.
//2. It must start with a letter.
//3. It can only contain letters, numbers, and the underscore character.
//4. It cannot end with an underscore character.
//
//If the username is valid then your program should return the string true, otherwise return the string false.

func CodelandUsernameValidation(str string) string {
	// sanity check
	if str == "" {
		return "" // do not do anything if the passed in value is an empty string
	}

	// Check the length of the username
	if len(str) < 4 || len(str) > 25 {
		return "false"
	}

	// Check if the username starts with a letter
	if !unicode.IsLetter(rune(str[0])) {
		return "false"
	}

	// Check if the username ends with an underscore
	if str[len(str)-1] == '_' {
		return "false"
	}

	// Check if the username contains only letters, numbers, and underscore
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return "false"
		}
	}

	return "true"
}
