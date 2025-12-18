package main

import "strings"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// compare lowercase versions
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[left])) {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= '0' && b <= '9') ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}
