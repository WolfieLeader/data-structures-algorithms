package matrix

import (
	"cmp"
	"fmt"
	"strings"
)

func (a *Matrix[T]) Get(row int, col int) (T, bool) {
	if row < 0 || row >= a.rows || col < 0 || col >= a.cols {
		var zero T
		return zero, false
	}
	return a.data[row][col], true
}

func (a *Matrix[T]) GetRow(row int) []T {
	if row < 0 || row >= a.rows {
		return nil
	}

	out := make([]T, a.cols)
	copy(out, a.data[row])
	return out
}

func (a *Matrix[T]) GetCol(col int) []T {
	if col < 0 || col >= a.cols {
		return nil
	}

	out := make([]T, a.rows)
	for i := 0; i < a.rows; i++ {
		out[i] = a.data[i][col]
	}
	return out
}

func (a *Matrix[T]) Set(row int, col int, value T) bool {
	if row < 0 || row >= a.rows || col < 0 || col >= a.cols {
		return false
	}

	a.data[row][col] = value
	return true
}

func (a *Matrix[T]) Replace(values ...[]T) bool {
	if len(values) > a.rows {
		return false
	}

	for _, row := range values {
		if len(row) > a.cols {
			return false
		}
	}

	var zero T
	a.Fill(zero)

	for i := 0; i < len(values); i++ {
		copy(a.data[i], values[i])
	}
	return true
}

func (a *Matrix[T]) Fill(value T) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = value
		}
	}
}

func (a *Matrix[T]) Clear() {
	var zero T
	a.Fill(zero)
}

func (a *Matrix[T]) Rows() int {
	return a.rows
}

func (a *Matrix[T]) Cols() int {
	return a.cols
}

func (a *Matrix[T]) Dimensions() (int, int) {
	return a.rows, a.cols
}

func (a *Matrix[T]) IsEmpty() bool {
	return a.rows == 0 || a.cols == 0
}

func (a *Matrix[T]) Search(value T) (int, int) {
	for row := 0; row < a.rows; row++ {
		for col := 0; col < a.cols; col++ {
			if cmp.Compare(a.data[row][col], value) == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func (a *Matrix[T]) Contains(value T) bool {
	row, col := a.Search(value)
	return row != -1 && col != -1
}

func (a *Matrix[T]) Traverse(fn func(row int, col int, value T)) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			fn(i, j, a.data[i][j])
		}
	}
}

func (a *Matrix[T]) Swap(row1, col1, row2, col2 int) bool {
	if row1 < 0 || row1 >= a.rows || col1 < 0 || col1 >= a.cols ||
		row2 < 0 || row2 >= a.rows || col2 < 0 || col2 >= a.cols {
		return false
	}

	a.data[row1][col1], a.data[row2][col2] = a.data[row2][col2], a.data[row1][col1]
	return true
}

func (a *Matrix[T]) SwapRow(row1, row2 int) bool {
	if row1 < 0 || row1 >= a.rows || row2 < 0 || row2 >= a.rows {
		return false
	}

	a.data[row1], a.data[row2] = a.data[row2], a.data[row1]
	return true
}

func (a *Matrix[T]) SwapCol(col1, col2 int) bool {
	if col1 < 0 || col1 >= a.cols || col2 < 0 || col2 >= a.cols {
		return false
	}

	for i := 0; i < a.rows; i++ {
		a.data[i][col1], a.data[i][col2] = a.data[i][col2], a.data[i][col1]
	}
	return true
}

func (a *Matrix[T]) Copy() Matrix[T] {
	out := make([][]T, a.rows)
	for i := 0; i < a.rows; i++ {
		out[i] = make([]T, a.cols)
		copy(out[i], a.data[i])
	}
	return Matrix[T]{data: out, rows: a.rows, cols: a.cols}
}

func (a *Matrix[T]) String() string {
	var sb strings.Builder
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			if j > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(fmt.Sprintf("%v", a.data[i][j]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
