package matrix

import "cmp"

type MatrixArray[T cmp.Ordered] interface {
	Replace(values ...[]T) error
	Get(row int, col int) (T, error)
	Set(row int, col int, value T) error
	Traverse()
	Length() int
	Clear()
	Copy() Matrix[T]
	IndexOf(value T) (int, int)
	Swap(index1, index2 [2]int) error
	Fill(value T)
}

var _ MatrixArray[int] = (*Matrix[int])(nil)
