package main

// TODO:

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int, len(nums)) // Number: Times
	for _, v := range nums {
		freqs[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqs {
		buckets[freq] = append(buckets[freq], num)
	}

	out := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0 && len(out) < k; i-- {
		for _, num := range buckets[i] {
			out = append(out, num)
			if len(out) == k {
				break
			}
		}
	}
	return out
}
