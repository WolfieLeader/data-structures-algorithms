package static

import "cmp"

type StaticArray[T cmp.Ordered] interface {
	Get(index int) (T, bool)
	Set(index int, value T) bool
	Replace(values ...T)
	Fill(value T)
	Clear()

	Length() int
	IsSorted() bool
	Search(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(index int, value T))
	Swap(index1, index2 int) bool

	ToSlice() []T
	Equal(other *Static[T]) bool
	Copy() *Static[T]
	Reverse() *Static[T]
	String() string
}

var _ StaticArray[int] = (*Static[int])(nil)
