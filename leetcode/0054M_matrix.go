package main

// TODO:

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
		for x := left; x <= right; x++ {
			out = append(out, matrix[top][x])
		}
		top++

		// NOTE: LOOP the right column
		for y := top; y <= bottom; y++ {
			out = append(out, matrix[y][right])
		}
		right--

		// Edge case single row
		if top <= bottom {
			// NOTE: LOOP the bottom row
			for x := right; x >= left; x-- {
				out = append(out, matrix[bottom][x])
			}
			bottom--
		}

		// Edge case single column
		if left <= right {
			// NOTE: LOOP the left column
			for y := bottom; y >= top; y-- {
				out = append(out, matrix[y][left])
			}
			left++
		}
	}

	return out
}
