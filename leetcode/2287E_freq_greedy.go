package main

func rearrangeCharacters(s string, target string) int {
	var freqSrc, freqTarget [26]int

	for i := 0; i < len(s); i++ {
		freqSrc[s[i]-'a']++
	}

	for i := 0; i < len(target); i++ {
		freqTarget[target[i]-'a']++
	}

	answer := 101 // Limit
	for i := range 26 {
		if freqTarget[i] > 0 {
			answer = min(answer, freqSrc[i]/freqTarget[i])
		}
	}

	return answer
}
