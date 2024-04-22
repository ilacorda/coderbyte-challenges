package main

import "unicode"

// Have the function SimpleSymbols(str) take the str parameter being passed and determine if it is an acceptable sequence by either returning the string true or false.
//The str parameter will be composed of + and = symbols with several characters between them (ie. ++d+===+c++==a) and
//for the string to be true each letter must be surrounded by a + symbol. So the string to the left would be false.
//The string will not be empty and will have at least one letter.

func SimpleSymbols(str string) bool {
	r := []rune(str)

	if len(str) == 0 {
		return false
	}

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(r[i]) {
			if i == 0 || i == len(runes)-1 || runes[i-1] != '+' || runes[i+1] != '+' {
				return false
			}
		}

	}

	return true

}
