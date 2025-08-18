package static

import "cmp"

type StaticArray[T cmp.Ordered] interface {
	Get(i int) (T, bool)
	Set(i int, value T) bool
	Replace(values ...T)
	Fill(value T)
	Clear()

	Length() int
	IsSorted() bool
	IsEmpty() bool
	Search(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(i int, value T) bool)
	Swap(i, j int) bool

	ToSlice() []T
	Copy() Static[T]
	Reverse() Static[T]
}

var _ StaticArray[int] = (*Static[int])(nil)
