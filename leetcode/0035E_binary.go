package main

// TODO

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) // [l, r) inclusive exclusive

	for left < right {
		mid := left + ((right - left) / 2)

		// Lower bound
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
