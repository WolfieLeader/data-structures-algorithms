package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] interface {
	Get(index int) (T, error)
	Set(index int, value T) error
	Replace(values ...T)
	Append(value ...T)
	Prepend(value ...T)
	Delete(index int) (T, error)
	Clear()
	Insert(index int, value T) error

	Length() int
	Capacity() int
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(func(index int, value T) bool)
	Swap(i, j int) error

	Slice(start, end int) (DynamicArray[T], error)
	Copy() DynamicArray[T]
	Reverse() DynamicArray[T]
}

var _ DynamicArray[int] = (*Dynamic[int])(nil)
