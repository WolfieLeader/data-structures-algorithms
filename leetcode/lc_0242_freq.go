package main

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
