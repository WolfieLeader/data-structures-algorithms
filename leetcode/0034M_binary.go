package main

// TODO

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 1) Find first index >= target (lower bound)
	l, r := 0, len(nums) // [l, r)
	for l < r {
		mid := l + ((r - l) / 2)

		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	start := l

	// If target is not present
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	// 2) Find first index > target (upper bound)
	l, r = 0, len(nums) // [l, r)
	for l < r {
		mid := l + ((r - l) / 2)
		if target >= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	end := l - 1

	return []int{start, end}
}
