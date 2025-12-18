package main

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
		if toLower(s[left]) != toLower(s[right]) {
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

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
