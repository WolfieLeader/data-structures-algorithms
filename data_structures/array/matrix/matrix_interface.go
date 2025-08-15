package matrix

import "cmp"

type MatrixArray[T cmp.Ordered] interface {
	Get(row int, col int) (T, error)
	GetRow(row int) []T
	GetCol(col int) []T
	Set(row int, col int, value T) error
	Replace(values ...[]T) error
	Fill(value T)
	Clear()

	Rows() int
	Cols() int
	Dims() (int, int)
	LinearSearch(value T) (int, int)
	Contains(value T) bool
	Traverse(func(row int, col int, value T) bool)
	Swap(row1, col1, row2, col2 int) error
	SwapRow(row1, row2 int) error
	SwapCol(col1, col2 int) error

	Copy() MatrixArray[T]
}

var _ MatrixArray[int] = (*Matrix[int])(nil)
