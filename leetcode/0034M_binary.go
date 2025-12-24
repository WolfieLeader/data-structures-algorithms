package main

// TODO

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 1) Find first index >= target (lower bound)
	left, right := 0, len(nums) // [l, r)
	for left < right {
		mid := left + ((right - left) / 2)

		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	start := left

	// If target is not present
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	// 2) Find first index > target (upper bound)
	left, right = 0, len(nums) // [l, r)
	for left < right {
		mid := left + ((right - left) / 2)
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	end := left - 1

	return []int{start, end}
}
