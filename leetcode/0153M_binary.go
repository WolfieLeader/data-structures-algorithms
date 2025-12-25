package main

func findMin(nums []int) int {
	l, r := 0, len(nums) // [l, r)
	last := nums[len(nums)-1]

	for l < r {
		mid := l + ((r - l) / 2)

		if nums[mid] <= last {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
