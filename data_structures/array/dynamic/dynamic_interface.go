package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] interface {
	Replace(values ...T)
	Get(index int) (T, error)
	Set(index int, value T) error
	Insert(index int, value T) error
	Append(value ...T)
	Prepend(value ...T)
	Delete(index int) (T, error)
	Length() int
	Capacity() int
	Traverse()
	Clear()
	Copy() Dynamic[T]
	Slice(start, end int) (Dynamic[T], error)
	IndexOf(value T) int
	Swap(i, j int) error
	Reverse() Dynamic[T]
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
}

var _ DynamicArray[int] = (*Dynamic[int])(nil)
