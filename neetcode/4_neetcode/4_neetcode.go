package main

type Freq [26]byte

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Freq][]string, len(strs))

	for _, str := range strs {
		var freq Freq
		for index := 0; index < len(str); index++ {
			freq[str[index]-'a']++
		}
		groups[freq] = append(groups[freq], str)
	}

	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}
