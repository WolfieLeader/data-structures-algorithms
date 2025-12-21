package main

// Sliding Window + Hash Map (Last Seen Index)
//
// Maintain a window [start..end] with all unique chars.
// A map stores the last index where each char was seen.
//
// If the current char was seen inside the window,
// move `start` to one position after its `last` index
// to ensure all characters remain unique.
//
// The window always represents a valid substring without repeats,
// and we track the maximum window length.
//
// Walkthrough: "pwwkew"
// 'p' - end=0, start=0,      -->  window="p",   len=1
// 'w' - end=1, start=0,      -->  window="pw",  len=2
// 'w' - end=2, start=0 -> 2  -->  window="w",   len=1
// 'k' - end=3, start=2       -->  window="wk",  len=2
// 'e' - end=4, start=2       -->  window="wke", len=3
// 'w' - end=5, start=2 -> 3  -->  window="kew", len=3

func lengthOfLongestSubstring(s string) int {
	seenAt := make(map[rune]int)

	best := 0
	start := 0
	for end, char := range s {
		if last, found := seenAt[char]; found && last >= start {
			start = last + 1
		}
		seenAt[char] = end

		windowSize := end - start + 1
		best = max(best, windowSize)
	}
	return best
}
