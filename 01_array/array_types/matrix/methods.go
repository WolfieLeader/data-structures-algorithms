package matrix

import (
	"cmp"
	"errors"
	"fmt"
)

type MatrixArray[T cmp.Ordered] interface {
	Replace(values ...[]T) error
	Get(row int, col int) (T, error)
	Set(row int, col int, value T) error
	Traverse()
	Length() int
	Clear()
	Copy() MatrixArrayType[T]
	IndexOf(value T) (int, int)
	Swap(index1, index2 [2]int) error
	Fill(value T)
}

var _ MatrixArray[int] = (*MatrixArrayType[int])(nil)

func (array *MatrixArrayType[T]) Replace(values ...[]T) error {
	if len(values) > array.rows {
		return errors.New("too many rows provided")
	}

	for i, row := range values {
		if len(row) > array.cols {
			return fmt.Errorf("row %d exceeds specified column count", i)
		}
	}

	for i := 0; i < array.rows; i++ {
		if i < len(values) {
			copy(array.data[i], values[i])
		} else {
			for j := range array.data[i] {
				array.data[i][j] = *new(T)
			}
		}
	}

	return nil
}

func (array *MatrixArrayType[T]) Get(row int, col int) (T, error) {
	if row < 0 || row >= array.rows || col < 0 || col >= array.cols {
		return *new(T), errors.New("index out of bounds")
	}
	return array.data[row][col], nil
}

func (array *MatrixArrayType[T]) Set(row int, col int, value T) error {
	if row < 0 || row >= array.rows || col < 0 || col >= array.cols {
		return errors.New("index out of bounds")
	}

	array.data[row][col] = value
	return nil
}

func (array *MatrixArrayType[T]) Traverse() {
	for i := range array.rows {
		for j := range array.cols {
			fmt.Printf("([%d][%d]:%v) ", i, j, array.data[i][j])
		}
		fmt.Println()
	}
}

func (array *MatrixArrayType[T]) Length() int {
	return array.rows * array.cols
}

func (array *MatrixArrayType[T]) Clear() {
	for i := range array.data {
		for j := range array.data[i] {
			array.data[i][j] = *new(T)
		}
	}
}

func (array *MatrixArrayType[T]) Copy() MatrixArrayType[T] {
	newData := make([][]T, array.rows)
	for i := range array.rows {
		newData[i] = make([]T, array.cols)
		copy(newData[i], array.data[i])
	}
	return MatrixArrayType[T]{data: newData, rows: array.rows, cols: array.cols}
}

func (array *MatrixArrayType[T]) IndexOf(value T) (int, int) {
	for i := range array.rows {
		for j := range array.cols {
			if cmp.Compare(array.data[i][j], value) == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (array *MatrixArrayType[T]) Swap(index1, index2 [2]int) error {
	i1, j1 := index1[0], index1[1]
	i2, j2 := index2[0], index2[1]

	if i1 < 0 || i1 >= array.rows || j1 < 0 || j1 >= array.cols ||
		i2 < 0 || i2 >= array.rows || j2 < 0 || j2 >= array.cols {
		return errors.New("index out of bounds")
	}

	array.data[i1][j1], array.data[i2][j2] = array.data[i2][j2], array.data[i1][j1]
	return nil
}

func (array *MatrixArrayType[T]) Fill(value T) {
	for i := range array.rows {
		for j := range array.cols {
			array.data[i][j] = value
		}
	}
}
