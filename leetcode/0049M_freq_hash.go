package main

func groupAnagrams(strs []string) [][]string {
	type Freq [26]int16
	groups := make(map[Freq][]string, len(strs))

	for _, str := range strs {
		var freq Freq
		for i := 0; i < len(str); i++ {
			freq[str[i]-'a']++
		}
		groups[freq] = append(groups[freq], str)
	}

	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}
