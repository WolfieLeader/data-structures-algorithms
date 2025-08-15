package static

import "cmp"

type StaticArray[T cmp.Ordered] interface {
	Get(i int) (T, error)
	Set(i int, value T) error
	Replace(values ...T)
	Fill(value T)
	Clear()

	Length() int
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(i int, value T) bool)
	Swap(i, j int) error

	ToSlice() []T
	Copy() Static[T]
	Reverse() Static[T]
}

var _ StaticArray[int] = (*Static[int])(nil)
