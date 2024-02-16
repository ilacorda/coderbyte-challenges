package main

func foo(param []int) int {
	if len(param) == 0 {
		return 0 // if it is an empty string
	}

	maxFound := param[0]
	maxEnd := param[0]
	for i := 1; i < len(param); i++ {
		maxEnd = maxBetweenNums(param[i], maxEnd+param[i])
		maxFound = maxBetweenNums(maxFound, maxEnd)

	}
	return maxFound
}

func maxBetweenNums(a, b int) int {
	if a > b {
		return a
	}
	return b
}
