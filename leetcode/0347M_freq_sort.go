package main

import "sort"

// TODO:

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, x := range nums {
		freq[x]++
	}

	uniq := make([]int, 0, len(freq))
	for x := range freq {
		uniq = append(uniq, x)
	}

	sort.Slice(uniq, func(i, j int) bool { return freq[uniq[i]] > freq[uniq[j]] })

	return uniq[:k]
}
