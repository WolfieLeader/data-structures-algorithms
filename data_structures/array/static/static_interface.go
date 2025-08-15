package static

import "cmp"

type StaticArray[T cmp.Ordered] interface {
	Replace(values ...T)
	Get(index int) (T, error)
	Set(index int, value T) error
	Length() int
	Traverse()
	Clear()
	Copy() Static[T]
	IndexOf(value T) int
	Swap(i, j int) error
	Reverse() Static[T]
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
}

var _ StaticArray[int] = (*Static[int])(nil)
