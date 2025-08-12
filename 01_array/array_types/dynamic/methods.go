package dynamic

import (
	"cmp"
	"errors"
	"fmt"
)

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
	Copy() DynamicArrayType[T]
	Slice(start, end int) (DynamicArrayType[T], error)
	IndexOf(value T) int
	Swap(i, j int) error
	Reverse() DynamicArrayType[T]
	IsSorted() bool
	LinearSearch(value T) int
	BinarySearch(value T) int
}

var _ DynamicArray[int] = (*DynamicArrayType[int])(nil)

func (array *DynamicArrayType[T]) Replace(values ...T) {
	*array = append((*array)[:0], values...)
}

func (array DynamicArrayType[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(array) {
		return *new(T), errors.New("index out of bounds")
	}
	return array[index], nil
}

func (array *DynamicArrayType[T]) Set(index int, value T) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*array)[index] = value
	return nil
}

func (array *DynamicArrayType[T]) Insert(index int, value T) error {
	if index < 0 || index > len(*array) {
		return errors.New("index out of bounds")
	}

	*array = append(*array, *new(T))           // grow by 1
	copy((*array)[index+1:], (*array)[index:]) // shift right
	(*array)[index] = value

	return nil
}

func (array *DynamicArrayType[T]) Append(value ...T) {
	*array = append(*array, value...)
}

func (array *DynamicArrayType[T]) Prepend(value ...T) {
	*array = append(value, *array...)
}

func (array *DynamicArrayType[T]) Delete(index int) (T, error) {
	if index < 0 || index >= len(*array) {
		return *new(T), errors.New("index out of bounds")
	}

	value := (*array)[index]
	*array = append((*array)[:index], (*array)[index+1:]...)
	return value, nil
}

func (array DynamicArrayType[T]) Length() int {
	return len(array)
}

func (array DynamicArrayType[T]) Capacity() int {
	return cap(array)
}

func (array DynamicArrayType[T]) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%v, ", i, value)
	}
	fmt.Println()
}

func (array *DynamicArrayType[T]) Clear() {
	*array = DynamicArrayType[T]{}
}

func (array DynamicArrayType[T]) Copy() DynamicArrayType[T] {
	newArray := make(DynamicArrayType[T], len(array))
	copy(newArray, array)
	return newArray
}

func (array DynamicArrayType[T]) Slice(start, end int) (DynamicArrayType[T], error) {
	if start < 0 || end > len(array) || start >= end {
		return nil, errors.New("invalid slice range")
	}
	slicedData := make(DynamicArrayType[T], end-start)
	copy(slicedData, array[start:end])
	return slicedData, nil
}

func (array DynamicArrayType[T]) IndexOf(value T) int {
	for i, v := range array {
		if cmp.Compare(v, value) == 0 {
			return i
		}
	}
	return -1
}

func (array *DynamicArrayType[T]) Swap(i, j int) error {
	if i < 0 || i >= len(*array) || j < 0 || j >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	return nil
}

func (array DynamicArrayType[T]) Reverse() DynamicArrayType[T] {
	reversed := array.Copy()
	for l, r := 0, len(reversed)-1; l < r; l, r = l+1, r-1 {
		reversed[l], reversed[r] = reversed[r], reversed[l]
	}
	return reversed
}

func (array DynamicArrayType[T]) IsSorted() bool {
	for i := 1; i < len(array); i++ {
		if cmp.Less(array[i], array[i-1]) {
			return false
		}
	}
	return true
}
