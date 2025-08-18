package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] interface {
	Get(i int) (T, bool)
	Set(i int, value T) bool
	Replace(values ...T)
	Append(values ...T)
	Prepend(values ...T)
	Insert(i int, value T) bool
	Delete(i int) (T, bool)
	Fill(value T)
	Clear()

	Length() int
	Capacity() int
	IsEmpty() bool
	IsSorted() bool
	Search(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(i int, value T) bool)
	Swap(i, j int) bool

	Slice(start, end int) (Dynamic[T], bool)
	Copy() Dynamic[T]
	Reverse() Dynamic[T]
}

var _ DynamicArray[int] = (*Dynamic[int])(nil)
