package main

// Sliding Window with frequency map
//
// Pattern
// Use a dynamic window [left..right] that expands to the right.
// Track how many times each character appears in the window.
//
// Invariant
// The window always contains only unique characters.
//
// When a duplicate appears
// Shrink the window from the left until the duplicate is removed.
// This restores the invariant.
//
// Why it works
// Each character enters and leaves the window at most once.
// This guarantees linear time.
//
// Walkthrough: "pwwkew"
// right=0 'p'  window="p"    left=0 best=1
// right=1 'w'  window="pw"   left=0 best=2
// right=2 'w'  duplicate -> shrink to window="w"   left=2
// right=3 'k'  window="wk"   left=2 best=2
// right=4 'e'  window="wke"  left=2 best=3
// right=5 'w'  duplicate -> shrink to window="kew" left=3

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)

	best := 0
	left := 0
	for right := 0; right < len(s); right++ {
		seen[s[right]]++

		// NOTE: LOOP shrink until char appears less than 1
		for seen[s[right]] > 1 {
			seen[s[left]]--
			left++
		}

		windowSize := right - left + 1
		best = max(best, windowSize)
	}
	return best
}

// Sliding Window + Hash Loop Up Optimization
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

func optimizedLengthOfLongestSubstring(s string) int {
	seenAt := make(map[rune]int)

	best := 0
	left := 0
	for right, char := range s {
		if last, found := seenAt[char]; found && last >= left {
			left = last + 1
		}
		seenAt[char] = right

		windowSize := right - left + 1
		best = max(best, windowSize)
	}
	return best
}
