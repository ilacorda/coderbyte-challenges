package main

import "testing"

func TestSearchingChallenge(t *testing.T) {

	testData := []struct {
		description string
		testStr     string
		expected    string
	}{
		{
			description: "it should return the longest palindromic substring",
			testStr:     "hellosannasmith",
			expected:    "sannas",
		},
		{
			description: "it should return none if no palindromic substring longer than 2 characters is found",
			testStr:     "abcd",
			expected:    "none",
		},
		{
			description: "it should return none no string is provided",
			testStr:     "",
			expected:    "none",
		},
	}

	for _, td := range testData {
		t.Run(td.description, func(t *testing.T) {
			actual := SearchingChallenge(td.testStr)
			if actual != td.expected {
				t.Errorf("Expected %s but got %s", td.expected, actual)
			}
		})
	}
}
