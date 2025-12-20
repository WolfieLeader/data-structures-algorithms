package main

//TODO:

// I: abcabcbb => Ex: abc/bca/cab length 3 => O: 3
// I: bbbb => 1
// I: pwwkew => 3

func lengthOfLongestSubstring(s string) int {
	seenAt := make(map[rune]int)
	longestLen := 0

	start := 0
	for end, ch := range []rune(s) {
		if index, found := seenAt[ch]; found && index >= start {
			start = index + 1
		}
		seenAt[ch] = end

		if winLen := (end - start + 1); winLen > longestLen {
			longestLen = winLen
		}
	}

	return longestLen
}
