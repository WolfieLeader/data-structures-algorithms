package main

// HACK: Basic Pattern of Two Pointer from Both Ends

// Two Pointers with filtering
//
// Pattern
// Use two pointers starting from both ends of the string.
// Move inward while maintaining a comparison invariant.
//
// Key idea
// Skip characters that should not participate in the comparison.
// Only compare valid characters.
//
// Invariant
// left and right always point to alphanumeric characters
// that should be compared.
//
// When pointers move
// If characters match, move both pointers inward.
// If they do not match, the string is not a palindrome.
//
// Why it works
// Each character is visited at most once.
// Skipping invalid characters does not affect correctness.
// Time complexity is linear.
//
// Walkthrough: "A man, a plan, a canal: Panama"
// left=A right=a match
// skip non alphanumeric characters
// continue inward until pointers cross
// result = true

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// NOTE: LOOP until left char is alphanumeric (stay in bounds!)
		for !isAlphanumeric(s[left]) && left < right {
			left++ // Forward
		}

		// NOTE: LOOP until right char is alphanumeric (stay in bounds!)
		for !isAlphanumeric(s[right]) && left < right {
			right-- // Backward
		}

		// Can use strings.ToLower(string(x)) instead
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
