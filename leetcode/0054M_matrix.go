package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	out := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// NOTE: LOOP the top row
		for c := left; c <= right; c++ {
			out = append(out, matrix[top][c])
		}
		top++

		// NOTE: LOOP the right column
		for r := top; r <= bottom; r++ {
			out = append(out, matrix[r][right])
		}
		right--

		// Edge case single row
		if top <= bottom {
			// NOTE: LOOP the bottom row
			for c := right; c >= left; c-- {
				out = append(out, matrix[bottom][c])
			}
			bottom--
		}

		// Edge case single column
		if left <= right {
			// NOTE: LOOP the left column
			for r := bottom; r >= top; r-- {
				out = append(out, matrix[r][left])
			}
			left++
		}
	}

	return out
}
