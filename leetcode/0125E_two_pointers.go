package main

// TODO:

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// NOTE: LOOP until left char is alphanumeric (stay in bounds!)
		for left < right && !isAlphanumeric(s[left]) {
			left++ // Forward
		}

		// NOTE: LOOP until right char is alphanumeric (stay in bounds!)
		for left < right && !isAlphanumeric(s[right]) {
			right-- // Backward
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++  // Forward
		right-- // Backward
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
