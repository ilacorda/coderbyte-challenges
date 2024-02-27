package main

import (
	"strconv"
)

func FizzBuzz(num int) string {
	result := ""

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			result += "FizzBuzz "
		} else if i%3 == 0 { // Check for numbers divisible by 3
			result += "Fizz "
		} else if i%5 == 0 { // Check for numbers divisible by 5
			result += "Buzz "
		} else { // For numbers not divisible by either 3 or 5, add the number itself
			result += strconv.Itoa(i) + " "
		}
	}

	return result[:len(result)-1] // Remove the trailing space
}
