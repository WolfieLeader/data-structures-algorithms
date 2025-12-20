package main

// TODO:

func partitionLabels(s string) []int {
	var lastIndex [26]int
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]-'a'] = i
	}

	out := make([]int, 0)
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		if lastIndex[s[i]-'a'] > end {
			end = lastIndex[s[i]-'a']
		}
		if i == end {
			out = append(out, end-start+1)
			start = i + 1
		}
	}

	return out
}
