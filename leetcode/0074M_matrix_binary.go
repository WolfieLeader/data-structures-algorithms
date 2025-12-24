package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m*n // [l, r)

	for left < right {
		mid := left + (right-left)/2

		row := mid / n
		col := mid % n
		val := matrix[row][col]

		if target == val {
			return true
		} else if target < val {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}
