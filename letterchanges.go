package main

import "strings"

//Letter Changes
//Have the function LetterChanges(str) take the str parameter being passed and modify it using the following algorithm.
//Replace every letter in the string with the letter following it in the alphabet (ie. c becomes d, z becomes a).
//Then capitalize every vowel in this new string (a, e, i, o, u) and finally return this modified string.

//Examples
//Input: "hello*3"
//Output: Ifmmp*3
//Input: "fun times!"
//Output: gvO Ujnft!

const englishAlphabet = "abcdefghijklmnopqrstuvwxyz"

func LetterChanges(str string) string {

	if len(str) == 0 {
		return ""
	}

	result := replaceLetter(str)

	capitalizedResultStr := capitalizeVowels(result)

	return capitalizedResultStr

}

func replaceLetter(str string) string {
	result := ""

	for _, char := range str {
		if index := strings.IndexRune(englishAlphabet, char); index != -1 {
			// Replace the letter with the next one
			nextIndex := (index + 1) % 26
			if char >= 'A' && char <= 'Z' {
				// Preserve uppercase letters
				result += strings.ToUpper(string(englishAlphabet[nextIndex]))
			} else {
				result += string(englishAlphabet[nextIndex])
			}
		} else {
			// Keep non-alphabetic characters unchanged
			result += string(char)
		}
	}

	return result
}

func capitalizeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			result += strings.ToUpper(string(char))
		} else {
			result += string(char)
		}
	}

	return result
}
