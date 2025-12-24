package main

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	// 1) Count frequencies
	for _, v := range arr {
		count[v]++
	}

	freqSeen := make(map[int]struct{})

	// 2) Check uniqueness of frequencies
	for _, freq := range count {
		if _, exists := freqSeen[freq]; exists {
			return false
		}
		freqSeen[freq] = struct{}{}
	}

	return true
}
