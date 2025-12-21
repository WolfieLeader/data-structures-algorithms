package main

// HACK: Core Pattern for Frequency

// Frequency Balance
//
// Pattern
// Count character frequencies using a fixed size array.
// Increment for one string and decrement for the other.
//
// Key idea
// Two strings are anagrams if all character counts balance to zero.
//
// Invariant
// After processing both strings,
// each character frequency represents the difference
// between its occurrences in the two strings.
//
// Early exit
// If lengths differ, they cannot be anagrams.
//
// Why it works
// Each character is processed once.
// Fixed alphabet size gives constant space.
// Time complexity is linear.
//
// Walkthrough: s="racecar", t="carrace"
// r +1 then -1 -> 0
// a +1 then -1 -> 0
// c +1 then -1 -> 0
// e +1 then -1 -> 0
// all frequencies are zero -> true

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int16
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for _, f := range freq {
		if f != 0 {
			return false
		}
	}
	return true
}
