package static

import (
	"cmp"
	"fmt"
)

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

func (array *Static[T]) Replace(values ...T) {
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy((*array)[:], values)
}

func (array Static[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(array) {
		return *new(T), fmt.Errorf("index out of bounds: %d", index)
	}

	return array[index], nil
}

func (array *Static[T]) Set(index int, value T) error {
	if index < 0 || index >= len(array) {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	(*array)[index] = value
	return nil
}

func (array Static[T]) Length() int {
	return len(array)
}

func (array Static[T]) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%v, ", i, value)
	}
	fmt.Println()
}

func (array *Static[T]) Clear() {
	*array = Static[T]{}
}

func (array Static[T]) Copy() Static[T] {
	var copyArray Static[T]
	copy(copyArray[:], array[:])
	return copyArray
}

func (array Static[T]) IndexOf(value T) int {
	for i, v := range array {
		if cmp.Compare(v, value) == 0 {
			return i
		}
	}
	return -1
}

func (array *Static[T]) Swap(i, j int) error {
	if i < 0 || i >= len(array) || j < 0 || j >= len(array) {
		return fmt.Errorf("index out of bounds: %d or %d", i, j)
	}
	array[i], array[j] = array[j], array[i]
	return nil
}

func (array Static[T]) Reverse() Static[T] {
	reversed := array.Copy()
	for l, r := 0, len(reversed)-1; l < r; l, r = l+1, r-1 {
		reversed[l], reversed[r] = reversed[r], reversed[l]
	}
	return reversed
}

func (array Static[T]) IsSorted() bool {
	for i := 1; i < len(array); i++ {
		if cmp.Less(array[i], array[i-1]) {
			return false
		}
	}
	return true
}
