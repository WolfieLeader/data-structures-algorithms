package main

func smallerNumbersThanCurrent(nums []int) []int {
	const maxVal = 101

	// 1) Frequency count
	count := make([]int, maxVal)
	for _, v := range nums {
		count[v]++
	}

	// 2) Prefix sum: count[i] = how many numbers < i
	for i := 1; i < maxVal; i++ {
		count[i] += count[i-1]
	}

	// 3) Build answer
	answer := make([]int, len(nums))
	for i, v := range nums {
		if v == 0 {
			answer[i] = 0
		} else {
			answer[i] = count[v-1]
		}
	}

	return answer
}
