package main

func isPerfectSquare(num int) bool {
	left, right := 1, num // [l, r]
	for left <= right {
		mid := left + ((right - left) / 2)
		squared := mid * mid

		if num == squared {
			return true
		} else if num < squared {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
