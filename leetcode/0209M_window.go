package main

// TODO:

// I: target = 7,  nums = [2,3,1,2,4,3] => O: 2
// I: target = 4,  nums = [1,4,4]  => O: 1
// I: target = 11, nums = [1,1,1,1,1,1,1,1]  => O: 0

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	best := 100001 // Limit is 10k

	left := 0
	for right, v := range nums {
		sum += v

		// NOTE: LOOP until sum is less than target
		for sum >= target {
			windowSize := right - left + 1
			best = min(best, windowSize)

			sum -= nums[left]
			left++
		}
	}

	if best == 100001 {
		return 0
	}
	return best
}
