package matrix

import (
	"cmp"
	"fmt"
)

func (arr Matrix[T]) Get(row int, col int) (T, error) {
	if row < 0 || row >= arr.rows || col < 0 || col >= arr.cols {
		var zero T
		return zero, fmt.Errorf("index out of bounds: row %d, col %d", row, col)
	}
	return arr.data[row][col], nil
}

func (arr Matrix[T]) GetRow(row int) []T {
	if row < 0 || row >= arr.rows {
		return nil
	}

	out := make([]T, arr.cols)
	copy(out, arr.data[row])
	return out
}

func (arr Matrix[T]) GetCol(col int) []T {
	if col < 0 || col >= arr.cols {
		return nil
	}

	out := make([]T, arr.rows)
	for i := 0; i < arr.rows; i++ {
		out[i] = arr.data[i][col]
	}
	return out
}

func (arr *Matrix[T]) Set(row int, col int, value T) error {
	if row < 0 || row >= arr.rows || col < 0 || col >= arr.cols {
		return fmt.Errorf("index out of bounds: row %d, col %d", row, col)
	}

	arr.data[row][col] = value
	return nil
}

func (arr *Matrix[T]) Replace(values ...[]T) error {
	if len(values) > arr.rows {
		return fmt.Errorf("too many rows provided, got %d, want %d", len(values), arr.rows)
	}

	for i, row := range values {
		if len(row) > arr.cols {
			return fmt.Errorf("row %d exceeds specified column count, got %d, want %d", i, len(row), arr.cols)
		}
	}

	var zero T
	arr.Fill(zero)

	for i := 0; i < len(values); i++ {
		copy(arr.data[i], values[i])
	}
	return nil
}

func (arr *Matrix[T]) Fill(value T) {
	for i := 0; i < arr.rows; i++ {
		for j := 0; j < arr.cols; j++ {
			arr.data[i][j] = value
		}
	}
}

func (arr *Matrix[T]) Clear() {
	var zero T
	arr.Fill(zero)
}

func (arr Matrix[T]) Rows() int {
	return arr.rows
}

func (arr Matrix[T]) Cols() int {
	return arr.cols
}

func (arr Matrix[T]) Dimensions() (int, int) {
	return arr.rows, arr.cols
}

func (arr Matrix[T]) Search(value T) (int, int) {
	for row := 0; row < arr.rows; row++ {
		for col := 0; col < arr.cols; col++ {
			if cmp.Compare(arr.data[row][col], value) == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func (arr Matrix[T]) Contains(value T) bool {
	row, col := arr.Search(value)
	return row != -1 && col != -1
}

func (arr Matrix[T]) Traverse(fn func(row int, col int, value T) bool) {
	for i := 0; i < arr.rows; i++ {
		for j := 0; j < arr.cols; j++ {
			if !fn(i, j, arr.data[i][j]) {
				return
			}
		}
	}
}

func (arr *Matrix[T]) Swap(row1, col1, row2, col2 int) error {
	if row1 < 0 || row1 >= arr.rows || col1 < 0 || col1 >= arr.cols ||
		row2 < 0 || row2 >= arr.rows || col2 < 0 || col2 >= arr.cols {
		return fmt.Errorf("index out of bounds: row1 %d, col1 %d, row2 %d, col2 %d", row1, col1, row2, col2)
	}

	arr.data[row1][col1], arr.data[row2][col2] = arr.data[row2][col2], arr.data[row1][col1]
	return nil
}

func (arr *Matrix[T]) SwapRow(row1, row2 int) error {
	if row1 < 0 || row1 >= arr.rows || row2 < 0 || row2 >= arr.rows {
		return fmt.Errorf("index out of bounds: row1 %d, row2 %d", row1, row2)
	}

	arr.data[row1], arr.data[row2] = arr.data[row2], arr.data[row1]
	return nil
}

func (arr *Matrix[T]) SwapCol(col1, col2 int) error {
	if col1 < 0 || col1 >= arr.cols || col2 < 0 || col2 >= arr.cols {
		return fmt.Errorf("index out of bounds: col1 %d, col2 %d", col1, col2)
	}

	for i := 0; i < arr.rows; i++ {
		arr.data[i][col1], arr.data[i][col2] = arr.data[i][col2], arr.data[i][col1]
	}
	return nil
}

func (arr Matrix[T]) Copy() Matrix[T] {
	out := make([][]T, arr.rows)
	for i := 0; i < arr.rows; i++ {
		out[i] = make([]T, arr.cols)
		copy(out[i], arr.data[i])
	}
	return Matrix[T]{data: out, rows: arr.rows, cols: arr.cols}
}
