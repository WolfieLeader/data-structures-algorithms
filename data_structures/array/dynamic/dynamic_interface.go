package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] interface {
	Get(index int) (T, bool)
	Set(index int, value T) bool
	Replace(values ...T)
	Append(values ...T)
	Prepend(values ...T)
	Insert(index int, value T) bool
	Delete(index int) (T, bool)
	Fill(value T)
	Clear()

	Length() int
	Capacity() int
	IsEmpty() bool
	IsSorted() bool
	Search(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(index int, value T))
	Swap(index1, index2 int) bool

	Between(start, end int) *Dynamic[T]
	Copy() *Dynamic[T]
	Reverse() *Dynamic[T]
	ToSlice() []T
	String() string
}

var _ DynamicArray[int] = (*Dynamic[int])(nil)
