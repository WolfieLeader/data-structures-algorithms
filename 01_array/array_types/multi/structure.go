package multi

import (
	"errors"
	"fmt"
)

type twoDimensionalArray struct {
	Data [][]int
	Rows int
	Cols int
}

func New2DArray(rows, cols int, values ...[]int) (twoDimensionalArray, error) {
	if len(values) > rows {
		return twoDimensionalArray{}, errors.New("too many rows provided")
	}

	for i, row := range values {
		if len(row) > cols {
			return twoDimensionalArray{}, fmt.Errorf("row %d exceeds specified column count", i)
		}
	}

	data := make([][]int, rows)
	for i := range rows {
		row := make([]int, cols)
		if i < len(values) {
			copy(row, values[i]) // prevent shared reference
		}
		data[i] = row
	}

	return twoDimensionalArray{Data: data, Rows: rows, Cols: cols}, nil
}
