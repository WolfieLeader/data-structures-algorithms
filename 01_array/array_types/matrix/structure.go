package matrix

import (
	"cmp"
	"errors"
	"fmt"
)

type MatrixArrayType[T cmp.Ordered] struct {
	data [][]T
	rows int
	cols int
}

func New[T cmp.Ordered](rows, cols int, values ...[]T) (MatrixArrayType[T], error) {
	if len(values) > rows {
		return MatrixArrayType[T]{}, errors.New("too many rows provided")
	}

	for i, row := range values {
		if len(row) > cols {
			return MatrixArrayType[T]{}, fmt.Errorf("row %d exceeds specified column count", i)
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

	return MatrixArrayType[T]{data: data, rows: rows, cols: cols}, nil
}
