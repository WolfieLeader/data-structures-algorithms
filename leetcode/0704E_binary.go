package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + ((right - left) / 2) // Against overflow
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
