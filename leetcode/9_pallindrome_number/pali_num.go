package main

// O(n) time | O(1) space
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	// Negative numbers and numbers ending with 0 (but not 0 itself) are not palindromes
	if x < 0 || x%10 == 0 {
		return false
	}

	num, reverse := x, 0
	for num != 0 {
		reverse = (reverse * 10) + (num % 10)
		num = num / 10
	}
	return x == reverse
}

func fasterIsPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	left, right := x, 0
	for left > right {
		right = (right * 10) + (left % 10)
		left = left / 10
	}

	// For odd digit numbers, we can get rid of the middle digit by right/10
	return left == right || left == right/10
}
