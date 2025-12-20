package main

//TODO:

// I: abcabcbb => Ex: abc/bca/cab length 3 => O: 3
// I: bbbb => 1
// I: pwwkew => 3

func lengthOfLongestSubstring(s string) int {
	seenAt := make(map[rune]int)
	longestLen := 0

	left := 0
	for right, ch := range []rune(s) {
		if i, found := seenAt[ch]; found && i >= left {
			left = i + 1
		}
		seenAt[ch] = right

		if winLen := (right - left + 1); winLen > longestLen {
			longestLen = winLen
		}
	}

	return longestLen
}
