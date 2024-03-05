package main

//Simple Adding
//Have the function SimpleAdding(num) add up all the numbers from 1 to num.
//For example: if the input is 4 then your program should return 10 because 1 + 2 + 3 + 4 = 10.

//Examples
//Input: 12
//Output: 78
//Input: 140
//Output: 9870

func SimpleAdding(num int) int {
	if num == 0 {
		return 0
	}

	result := 0

	for i := 1; i <= num; i++ {
		result += i

	}

	return result

}
