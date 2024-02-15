package main

import (
	"regexp"
	"strings"
)

func LongestWord(sen string) string {
	if len(sen) == 0 {
		return "length zero"
	}

	reg := regexp.MustCompile(`[.,!?;:'"&()[\]{}\-_]`)

	senWithNoPunctutation := reg.ReplaceAllString(sen, "")

	words := strings.Split(senWithNoPunctutation, " ")

	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}

	}

	return longestWord

}
