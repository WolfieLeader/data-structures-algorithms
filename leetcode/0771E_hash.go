package main

func numJewelsInStones(jewels string, stones string) int {
	jewelsSet := make(map[byte]struct{})
	for i := 0; i < len(jewels); i++ {
		jewelsSet[jewels[i]] = struct{}{}
	}

	count := 0
	for i := 0; i < len(stones); i++ {
		if _, found := jewelsSet[stones[i]]; found {
			count++
		}
	}
	return count
}
