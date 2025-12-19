package main

// I: leetcode => O: 0
// I: loveleetcode => O: 2
// I: aabb => O: -1

func firstUniqChar(s string) int {
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
