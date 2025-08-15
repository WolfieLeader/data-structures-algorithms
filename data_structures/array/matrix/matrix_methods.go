package matrix

import (
	"cmp"
	"fmt"
)

func (a matrixArray[T]) Get(row int, col int) (T, error) {
	if row < 0 || row >= a.rows || col < 0 || col >= a.cols {
		var zeroValue T
		return zeroValue, fmt.Errorf("index out of bounds: row %d, col %d", row, col)
	}
	return a.data[row][col], nil
}

func (a matrixArray[T]) GetRow(row int) []T {
	if row < 0 || row >= a.rows {
		return nil
	}

	out := make([]T, a.cols)
	copy(out, a.data[row])
	return out
}

func (a matrixArray[T]) GetCol(col int) []T {
	if col < 0 || col >= a.cols {
		return nil
	}

	out := make([]T, a.rows)
	for i := 0; i < a.rows; i++ {
		out[i] = a.data[i][col]
	}
	return out
}

func (a *matrixArray[T]) Set(row int, col int, value T) error {
	if row < 0 || row >= (*a).rows || col < 0 || col >= (*a).cols {
		return fmt.Errorf("index out of bounds: row %d, col %d", row, col)
	}

	(*a).data[row][col] = value
	return nil
}

func (a *matrixArray[T]) Replace(values ...[]T) error {
	if len(values) > (*a).rows {
		return fmt.Errorf("too many rows provided, got %d, want %d", len(values), a.rows)
	}

	for i, row := range values {
		if len(row) > (*a).cols {
			return fmt.Errorf("row %d exceeds specified column count, got %d, want %d", i, len(row), a.cols)
		}
	}

	var zeroValue T

	for i := 0; i < (*a).rows; i++ {
		if i < len(values) {
			copy((*a).data[i], values[i])
		} else {
			for j := range (*a).data[i] {
				(*a).data[i][j] = zeroValue
			}
		}
	}

	return nil
}

func (a *matrixArray[T]) Fill(value T) {
	for i := 0; i < (*a).rows; i++ {
		for j := 0; j < (*a).cols; j++ {
			(*a).data[i][j] = value
		}
	}
}

func (a *matrixArray[T]) Clear() {
	var zeroValue T
	a.Fill(zeroValue)
}

func (a matrixArray[T]) Rows() int {
	return a.rows
}

func (a matrixArray[T]) Cols() int {
	return a.cols
}

func (a matrixArray[T]) Dimensions() (int, int) {
	return a.rows, a.cols
}

func (a matrixArray[T]) LinearSearch(value T) (int, int) {
	for row := range a.rows {
		for col := range a.cols {
			if cmp.Compare(a.data[row][col], value) == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func (a matrixArray[T]) Contains(value T) bool {
	row, col := a.LinearSearch(value)
	return row != -1 && col != -1
}

func (a matrixArray[T]) Traverse(fn func(row int, col int, value T) bool) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			if !fn(i, j, a.data[i][j]) {
				return
			}
		}
	}
}

func (a *matrixArray[T]) Swap(row1, col1, row2, col2 int) error {
	if row1 < 0 || row1 >= (*a).rows || col1 < 0 || col1 >= (*a).cols ||
		row2 < 0 || row2 >= (*a).rows || col2 < 0 || col2 >= (*a).cols {
		return fmt.Errorf("index out of bounds: row1 %d, col1 %d, row2 %d, col2 %d", row1, col1, row2, col2)
	}

	(*a).data[row1][col1], (*a).data[row2][col2] = (*a).data[row2][col2], (*a).data[row1][col1]
	return nil
}

func (a *matrixArray[T]) SwapRow(row1, row2 int) error {
	if row1 < 0 || row1 >= (*a).rows || row2 < 0 || row2 >= (*a).rows {
		return fmt.Errorf("index out of bounds: row1 %d, row2 %d", row1, row2)
	}

	(*a).data[row1], (*a).data[row2] = (*a).data[row2], (*a).data[row1]
	return nil
}

func (a *matrixArray[T]) SwapCol(col1, col2 int) error {
	if col1 < 0 || col1 >= (*a).cols || col2 < 0 || col2 >= (*a).cols {
		return fmt.Errorf("index out of bounds: col1 %d, col2 %d", col1, col2)
	}

	for i := 0; i < (*a).rows; i++ {
		(*a).data[i][col1], (*a).data[i][col2] = (*a).data[i][col2], (*a).data[i][col1]
	}
	return nil
}

func (a matrixArray[T]) Copy() matrixArray[T] {
	out := make([][]T, a.rows)
	for i := range a.rows {
		out[i] = make([]T, a.cols)
		copy(out[i], a.data[i])
	}
	return matrixArray[T]{data: out, rows: a.rows, cols: a.cols}
}
