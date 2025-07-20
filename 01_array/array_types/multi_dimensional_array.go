package array_types

import (
	"errors"
	"fmt"
)

type twoDimensionalArray struct {
	Data [][]int
	Rows int
	Cols int
}

type TwoDimensionalArray interface {
	Replace(values ...[]int) error
	Get(row, col int) (int, error)
	Set(row, col, value int) error
	Traverse()
	Length() int
	Clear()
	Copy() twoDimensionalArray
	IndexOf(value int) (int, int)
	Swap(index1, index2 [2]int) error
	Fill(value int)
}

var _ TwoDimensionalArray = (*twoDimensionalArray)(nil)

func NewTwoDimensionalArray(rows, cols int, values ...[]int) (twoDimensionalArray, error) {
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

func (array *twoDimensionalArray) Replace(values ...[]int) error {
	if len(values) > array.Rows {
		return errors.New("too many rows provided")
	}

	for i, row := range values {
		if len(row) > array.Cols {
			return fmt.Errorf("row %d exceeds specified column count", i)
		}
	}

	for i := 0; i < array.Rows; i++ {
		if i < len(values) {
			copy(array.Data[i], values[i])
		} else {
			for j := range array.Data[i] {
				array.Data[i][j] = 0
			}
		}
	}

	return nil
}

func (array *twoDimensionalArray) Get(row, col int) (int, error) {
	if row < 0 || row >= array.Rows || col < 0 || col >= array.Cols {
		return 0, errors.New("index out of bounds")
	}
	return array.Data[row][col], nil
}

func (array *twoDimensionalArray) Set(row, col, value int) error {
	if row < 0 || row >= array.Rows || col < 0 || col >= array.Cols {
		return errors.New("index out of bounds")
	}
	array.Data[row][col] = value
	return nil
}

func (array *twoDimensionalArray) Traverse() {
	for i := 0; i < array.Rows; i++ {
		for j := 0; j < array.Cols; j++ {
			fmt.Printf("([%d][%d]:%d) ", i, j, array.Data[i][j])
		}
		fmt.Println()
	}
}

func (array *twoDimensionalArray) Length() int {
	return array.Rows * array.Cols
}

func (array *twoDimensionalArray) Clear() {
	for i := range array.Data {
		for j := range array.Data[i] {
			array.Data[i][j] = 0
		}
	}
}

func (array *twoDimensionalArray) Copy() twoDimensionalArray {
	newData := make([][]int, array.Rows)
	for i := 0; i < array.Rows; i++ {
		newData[i] = make([]int, array.Cols)
		copy(newData[i], array.Data[i])
	}
	return twoDimensionalArray{Data: newData, Rows: array.Rows, Cols: array.Cols}
}

func (array *twoDimensionalArray) IndexOf(value int) (int, int) {
	for i := 0; i < array.Rows; i++ {
		for j := 0; j < array.Cols; j++ {
			if array.Data[i][j] == value {
				return i, j
			}
		}
	}
	return -1, -1
}

func (array *twoDimensionalArray) Swap(index1, index2 [2]int) error {
	i1, j1 := index1[0], index1[1]
	i2, j2 := index2[0], index2[1]

	if i1 < 0 || i1 >= array.Rows || j1 < 0 || j1 >= array.Cols ||
		i2 < 0 || i2 >= array.Rows || j2 < 0 || j2 >= array.Cols {
		return errors.New("index out of bounds")
	}

	array.Data[i1][j1], array.Data[i2][j2] = array.Data[i2][j2], array.Data[i1][j1]
	return nil
}

func (array *twoDimensionalArray) Fill(value int) {
	for i := 0; i < array.Rows; i++ {
		for j := 0; j < array.Cols; j++ {
			array.Data[i][j] = value
		}
	}
}
