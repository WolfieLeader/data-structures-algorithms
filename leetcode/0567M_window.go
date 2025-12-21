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

	left := 0
	for right := 0; right < len(s2); right++ {
		freq[s2[right]-'a']--

		// NOTE: LOOP shrink if found letter not inside s1 or removed too many
		for freq[s2[right]-'a'] < 0 {
			freq[s2[left]-'a']++
			left++
		}

		if windowSize := (right - left + 1); windowSize == len(s1) {
			return true
		}
	}

	return false
}
