package main

func findDisappearedNumbers(nums []int) []int {
	// 1) Mark seen numbers by flipping sign at their mapped index.
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] *= -1
		}
	}

	// 2) Collect missing numbers (indexes still positive).
	missing := make([]int, 0)
	for i, v := range nums {
		if v > 0 {
			missing = append(missing, i+1)
		}
	}
	return missing
}
