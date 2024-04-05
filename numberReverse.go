package main

import "strings"

//Number Reverse
//Have the function NumberReverse(str) take the str parameter being passed which will be a string of numbers, and return a new string with the numbers in reverse order.
//Examples
//Input: "1 2 3"
//Output: 3 2 1
//Input: "10 20 50"
//Output: 50 20 10

func NumberReverse(str string) string {
	// sanitiy check
	if len(str) == 0 {
		return ""
	}

	rs := strings.Split(str, " ")
	result := ""

	for i := len(rs) - 1; i >= 0; i-- {
		result += rs[i] + " "
	}

	return result

}
