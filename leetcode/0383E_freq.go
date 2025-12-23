package main

func canConstruct(ransomNote string, magazine string) bool {
	var freq [26]int

	// Count letters in magazine
	for i := 0; i < len(magazine); i++ {
		freq[magazine[i]-'a']++
	}

	// Consume letters for ransomNote
	for i := 0; i < len(ransomNote); i++ {
		freq[ransomNote[i]-'a']--
		if freq[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
