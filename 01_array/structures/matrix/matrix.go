package matrix

import (
	"cmp"
	"errors"
	"fmt"
)

type Matrix[T cmp.Ordered] struct {
	data [][]T
	rows int
	cols int
}

func New[T cmp.Ordered](rows, cols int, values ...[]T) (Matrix[T], error) {
	if len(values) > rows {
		return Matrix[T]{}, errors.New("too many rows provided")
	}

	for i, row := range values {
		if len(row) > cols {
			return Matrix[T]{}, fmt.Errorf("row %d exceeds specified column count", i)
		}
	}

	data := make([][]T, rows)
	for i := range rows {
		row := make([]T, cols)
		if i < len(values) {
			copy(row, values[i]) // prevent shared reference
		}
		data[i] = row
	}

	return Matrix[T]{data: data, rows: rows, cols: cols}, nil
}
