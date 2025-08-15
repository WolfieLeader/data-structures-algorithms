package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] interface {
	Get(i int) (T, error)
	Set(i int, value T) error
	Replace(values ...T)
	Append(values ...T)
	Prepend(values ...T)
	Insert(i int, value T) error
	Delete(i int) (T, error)
	Fill(value T)
	Clear()

	Length() int
	Capacity() int
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(i int, value T) bool)
	Swap(i, j int) error

	Slice(start, end int) (dynamicArray[T], error)
	Copy() dynamicArray[T]
	Reverse() dynamicArray[T]
}

var _ DynamicArray[int] = (*dynamicArray[int])(nil)
