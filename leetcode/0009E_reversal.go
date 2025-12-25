package main

func isPalindrome9(x int) bool {
	if x == 0 {
		return true
	}
	// Negative numbers and numbers ending with 0 (but not 0 itself) are not palindromes
	if x < 0 || x%10 == 0 {
		return false
	}

	num, reverse := x, 0
	for num != 0 {
		reverse *= 10
		reverse += num % 10

		num /= 10
	}

	return x == reverse
}
