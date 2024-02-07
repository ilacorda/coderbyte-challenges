package main

func SearchingChallenge(str string) string {
	if len(str) == 0 {
		return "none"
	}

	longestPalindrome := ""

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(str) && str[left] == str[right] {
			palindrome := str[left : right+1]
			if len(palindrome) > len(longestPalindrome) {
				longestPalindrome = palindrome
			}
			left--
			right++
		}
	}

	for i := 0; i < len(str); i++ {
		// Check for odd-length palindromes
		expandAroundCenter(i, i)
		// Check for even-length palindromes
		expandAroundCenter(i, i+1)
	}

	// Return "none" if no palindromic substring longer than 2 characters is found
	if len(longestPalindrome) <= 2 {
		return "none"
	}
	return longestPalindrome
}
