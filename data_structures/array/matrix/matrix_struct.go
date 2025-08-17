package matrix

import (
	"cmp"
	"fmt"
)

type matrix[T cmp.Ordered] struct {
	data [][]T
	rows int
	cols int
}

func New[T cmp.Ordered](rows, cols int) matrix[T] {
	if rows < 0 || cols < 0 {
		panic(fmt.Sprintf("negative dimensions: rows=%d cols=%d", rows, cols))
	}
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	return matrix[T]{data: data, rows: rows, cols: cols}
}

func NewFromValues[T cmp.Ordered](values ...[]T) (matrix[T], error) {
	if len(values) == 0 {
		return matrix[T]{}, fmt.Errorf("no rows provided")
	}
	rows, cols := len(values), len(values[0])

	data := make([][]T, rows)
	for i, row := range values {
		if len(row) != cols {
			return matrix[T]{}, fmt.Errorf("ragged rows: row %d len=%d, want %d", i, len(row), cols)
		}
		cp := make([]T, len(row))
		copy(cp, row)
		data[i] = cp
	}
	return matrix[T]{data: data, rows: rows, cols: cols}, nil
}
