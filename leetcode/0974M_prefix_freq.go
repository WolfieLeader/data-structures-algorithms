package main

// TODO:

func subarraysDivByK(nums []int, k int) int {
	sumFreq := make([]int, k)
	sumFreq[0] = 1

	sum, answer := 0, 0
	for _, v := range nums {
		sum += v

		remainder := sum % k
		if remainder < 0 {
			remainder += k
		}

		answer += sumFreq[remainder]
		sumFreq[remainder]++
	}

	return answer
}
