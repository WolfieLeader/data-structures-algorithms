package main

// TODO:

// I: nums = [1,1,1], k = 2 => O: 2
// I: nums = [1,2,3], k = 3 => O: 2

func subarraySum(nums []int, k int) int {
	prefixFreqs := make(map[int]int) // Prefix Sum: Freq
	prefixFreqs[0] = 1

	count, sum := 0, 0
	for _, v := range nums {
		sum += v

		target := sum - k
		if freq, found := prefixFreqs[target]; found {
			count += freq
		}

		prefixFreqs[sum]++
	}

	return count
}
