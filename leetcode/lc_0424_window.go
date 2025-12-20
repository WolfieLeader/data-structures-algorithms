package main

//TODO:

// I: s = "ABAB", k = 2 => Ex: AAAA/BBBB len 4 => O: 4
// I: s = "AABABBA", k = 1 => O: 4

func characterReplacement(s string, k int) int {
	var freq [26]int
	best, maxFreq := 0, 0

	left := 0
	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		maxFreq = max(maxFreq, freq[s[right]-'A'])

		// if replacements needed > k shrink window
		for (right-left+1)-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}

		if winLen := (right - left + 1); winLen > best {
			best = winLen
		}
	}

	return best
}
