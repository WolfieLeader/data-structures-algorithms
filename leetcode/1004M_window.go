package main

func longestOnes(nums []int, k int) int {
	zeros := 0
	best := 0

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		// NOTE: LOOP until there is enough 1 in the window
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		windowSize := right - left + 1
		best = max(best, windowSize)
	}

	return best
}
