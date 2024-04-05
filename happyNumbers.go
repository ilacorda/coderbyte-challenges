package main

//Happy Numbers
//Have the function HappyNumbers(num) determine if a number is Happy, which is a number whose sum of the square of the digits eventually converges to 1. Return true if it's a Happy number, otherwise return false.
//
//For example: the number 10 is Happy because 1^2 + 0^2 converges to 1.
//Examples
//Input: 1
//Output: true
//Input: 101
//Output: false

func HappyNumbers(num int) bool {
	if num == 0 {
		return false
	}

	getNext := func(number int) int {
		sum := 0
		for number > 0 {
			digit := number % 10
			sum += digit * digit
			number /= 10
		}
		return sum
	}

	seen := make(map[int]struct{})
	for num != 1 {
		if _, ok := seen[num]; ok {
			// We've seen this number before, so it's not a happy number
			return false
		}
		seen[num] = struct{}{}
		num = getNext(num)
	}

	return true

}
