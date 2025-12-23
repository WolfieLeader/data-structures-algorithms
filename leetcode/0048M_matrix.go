package main

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	// 1) Transpose (rotate left)
	// [1 2 3]      [1 4 7]
	// [4 5 6]  ->  [2 5 8]
	// [7 8 9]      [3 6 9]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2) Reverse each row
	// [1 4 7]      [7 4 1]
	// [2 5 8]  ->  [8 5 2]
	// [3 6 9]      [9 6 3]
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
