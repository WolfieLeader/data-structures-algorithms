package matrix

import (
	"cmp"
	"fmt"
)

type Matrix[T cmp.Ordered] struct {
	data [][]T
	rows int
	cols int
}

func New[T cmp.Ordered](rows, cols int) Matrix[T] {
	if rows < 0 || cols < 0 {
		panic(fmt.Sprintf("negative dimensions: rows=%d cols=%d", rows, cols))
	}
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	return Matrix[T]{data: data, rows: rows, cols: cols}
}

func NewFromValues[T cmp.Ordered](values ...[]T) (Matrix[T], error) {
	if len(values) == 0 {
		return Matrix[T]{}, fmt.Errorf("no rows provided")
	}
	rows, cols := len(values), len(values[0])

	data := make([][]T, rows)
	for i, row := range values {
		if len(row) != cols {
			return Matrix[T]{}, fmt.Errorf("ragged rows: row %d len=%d, want %d", i, len(row), cols)
		}
		cp := make([]T, len(row))
		copy(cp, row)
		data[i] = cp
	}
	return Matrix[T]{data: data, rows: rows, cols: cols}, nil
}
