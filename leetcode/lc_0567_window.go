package main

// I: s1="ba", s2="fsdbafuia" => O: true

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var freq [26]int
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}

	start := 0
	for end := 0; end < len(s2); end++ {
		freq[s2[end]-'a']--

		// If we found char not in s1 or two many of the char from s1
		for freq[s2[end]-'a'] < 0 {
			freq[s2[start]-'a']++
			start++
		}

		if end-start+1 == len(s1) {
			return true
		}
	}

	return false
}
