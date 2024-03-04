package main

import "strconv"

//Questions Marks
//Have the function QuestionsMarks(str) take the str string parameter, which will contain single digit numbers, letters, and question marks,
//and check if there are exactly 3 question marks between every pair of two numbers that add up to 10.
//If so, then your program should return the string true, otherwise it should return the string false.
//If there aren't any two numbers that add up to 10 in the string, then your program should return false as well.
//
//For example: if str is "arrb6???4xxbl5???eee5" then your program should return true because there are exactly 3 question marks between 6 and 4, and 3 question marks between 5 and 5 at the end of the string.
//
//Examples
//Input: "aa6?9"
//Output: false
//Input: "acc?7??sss?3rr1??????5"
//Output: true

func QuestionsMarks(str string) string {
	lastNum := -1
	lastPos := -1
	found := false

	for i, char := range str {
		if isDigit(char) {
			num, _ := strconv.Atoi(string(char))
			if lastNum != -1 && (num+lastNum == 10) {
				if countQuestionMarks(str[lastPos+1:i]) == 3 {
					found = true
				} else {
					return "false"
				}
			}
			lastNum = num
			lastPos = i
		}
	}

	if found {
		return "true"
	}
	return "false"
}

func countQuestionMarks(s string) int {
	count := 0
	for _, char := range s {
		if char == '?' {
			count++
		}
	}
	return count

}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
