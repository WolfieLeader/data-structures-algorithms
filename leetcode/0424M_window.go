package main

//TODO:

// I: s = "ABAB", k = 2 => Ex: AAAA/BBBB len 4 => O: 4
// I: s = "AABABBA", k = 1 => O: 4

func characterReplacement(s string, k int) int {
	var freq [26]int
	best, highestCount := 0, 0

	left := 0
	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		highestCount = max(highestCount, freq[s[right]-'A'])

		// NOTE: LOOP shrink until `replacements <= k`
		for (right-left+1)-highestCount > k {
			freq[s[left]-'A']--
			left++
		}

		windowSize := right - left + 1
		best = max(best, windowSize)
	}

	return best
}
