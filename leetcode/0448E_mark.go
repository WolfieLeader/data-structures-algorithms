package main

import "math"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)

	// 1) Mark seen numbers by flipping sign at their mapped index.
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] > 0 {
			nums[index] *= (-1)
		}
	}

	// 2) Collect missing numbers (indexes still positive).
	missing := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			missing = append(missing, i+1)
		}
	}
	return missing
}
