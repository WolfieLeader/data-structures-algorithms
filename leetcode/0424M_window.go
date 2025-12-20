package main

//TODO:

// I: s = "ABAB", k = 2 => Ex: AAAA/BBBB len 4 => O: 4
// I: s = "AABABBA", k = 1 => O: 4

func characterReplacement(s string, k int) int {
	var freq [26]int
	best, maxFreq := 0, 0

	start := 0
	for end := 0; end < len(s); end++ {
		freq[s[end]-'A']++
		maxFreq = max(maxFreq, freq[s[end]-'A'])

		// if replacements needed > k shrink window
		for (end-start+1)-maxFreq > k {
			freq[s[start]-'A']--
			start++
		}

		if winLen := (end - start + 1); winLen > best {
			best = winLen
		}
	}

	return best
}
