package main

// TODO

func maxNumberOfBalloons(text string) int {
	var freqSrc, freqTarget [26]int

	for i := 0; i < len(text); i++ {
		freqSrc[text[i]-'a']++
	}

	for _, char := range "balloon" {
		freqTarget[char-'a']++
	}

	answer := 10001 // Limit
	for i := range 26 {
		if freqTarget[i] > 0 {
			answer = min(answer, freqSrc[i]/freqTarget[i])
		}
	}

	return answer
}
