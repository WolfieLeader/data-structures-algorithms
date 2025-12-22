package main

func findClosestNumber(nums []int) int {
	best := nums[0]

	for _, v := range nums {
		if abs(v) < abs(best) || (abs(v) == abs(best) && v > best) {
			best = v
		}
	}

	return best
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
