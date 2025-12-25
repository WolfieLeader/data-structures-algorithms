package main

// TODO

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums) // [l, r) inclusive exclusive

	for l < r {
		mid := l + ((r - l) / 2)

		// Lower bound
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
