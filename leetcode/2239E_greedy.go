package main

func findClosestNumber(nums []int) int {
	closeToZero := 100001 // Limit is 10K
	for _, v := range nums {
		if abs(v) < abs(closeToZero) || abs(v) == abs(closeToZero) && closeToZero < 0 && v > 0 {
			closeToZero = v
		}
	}
	return closeToZero
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
