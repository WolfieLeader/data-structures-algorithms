package main

// TODO:

// I: target = 7,  nums = [2,3,1,2,4,3] => O: 2
// I: target = 4,  nums = [1,4,4]  => O: 1
// I: target = 11, nums = [1,1,1,1,1,1,1,1]  => O: 0

func minSubArrayLen(target int, nums []int) int {
	sum, minLen := 0, 100001
	start := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			minLen = min(minLen, end-start+1)
			sum = sum - nums[start]
			start++
		}
	}

	if minLen == 100001 {
		return 0
	}
	return minLen
}
