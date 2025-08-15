package matrix

import (
	"cmp"
	"fmt"
)

type matrixArray[T cmp.Ordered] struct {
	data [][]T
	rows int
	cols int
}

func New[T cmp.Ordered](rows, cols int) matrixArray[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	return matrixArray[T]{data: data, rows: rows, cols: cols}
}

func NewFromValues[T cmp.Ordered](values ...[]T) (matrixArray[T], error) {
	if len(values) == 0 {
		return matrixArray[T]{}, fmt.Errorf("no rows provided")
	}

	rows := len(values)
	cols := len(values[0])
	data := make([][]T, rows)

	for i, row := range values {
		if len(row) != cols {
			return matrixArray[T]{}, fmt.Errorf("ragged rows")
		}
		copyRow := append([]T(nil), row...) // copy to avoid aliasing
		data[i] = copyRow
	}

	return matrixArray[T]{data: data, rows: rows, cols: cols}, nil
}
