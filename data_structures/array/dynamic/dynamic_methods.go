package dynamic

import (
	"cmp"
	"errors"
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
)

func (array *Dynamic[T]) Replace(values ...T) {
	*array = append((*array)[:0], values...)
}

func (array Dynamic[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(array) {
		return *new(T), errors.New("index out of bounds")
	}
	return array[index], nil
}

func (array *Dynamic[T]) Set(index int, value T) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*array)[index] = value
	return nil
}

func (array *Dynamic[T]) Insert(index int, value T) error {
	if index < 0 || index > len(*array) {
		return errors.New("index out of bounds")
	}

	*array = append(*array, *new(T))           // grow by 1
	copy((*array)[index+1:], (*array)[index:]) // shift right
	(*array)[index] = value

	return nil
}

func (array *Dynamic[T]) Append(value ...T) {
	*array = append(*array, value...)
}

func (array *Dynamic[T]) Prepend(value ...T) {
	*array = append(value, *array...)
}

func (array *Dynamic[T]) Delete(index int) (T, error) {
	if index < 0 || index >= len(*array) {
		return *new(T), errors.New("index out of bounds")
	}

	value := (*array)[index]
	*array = append((*array)[:index], (*array)[index+1:]...)
	return value, nil
}

func (array Dynamic[T]) Length() int {
	return len(array)
}

func (array Dynamic[T]) Capacity() int {
	return cap(array)
}

func (array Dynamic[T]) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%v, ", i, value)
	}
	fmt.Println()
}

func (array *Dynamic[T]) Clear() {
	*array = Dynamic[T]{}
}

func (array Dynamic[T]) Copy() Dynamic[T] {
	newArray := make(Dynamic[T], len(array))
	copy(newArray, array)
	return newArray
}

func (array Dynamic[T]) Slice(start, end int) (Dynamic[T], error) {
	if start < 0 || end > len(array) || start >= end {
		return nil, errors.New("invalid slice range")
	}
	slicedData := make(Dynamic[T], end-start)
	copy(slicedData, array[start:end])
	return slicedData, nil
}

func (array Dynamic[T]) IndexOf(value T) int {
	for i, v := range array {
		if cmp.Compare(v, value) == 0 {
			return i
		}
	}
	return -1
}

func (array *Dynamic[T]) Swap(i, j int) error {
	if i < 0 || i >= len(*array) || j < 0 || j >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	return nil
}

func (array Dynamic[T]) Reverse() Dynamic[T] {
	reversed := array.Copy()
	for l, r := 0, len(reversed)-1; l < r; l, r = l+1, r-1 {
		reversed[l], reversed[r] = reversed[r], reversed[l]
	}
	return reversed
}

func (array Dynamic[T]) IsSorted() bool {
	for i := 1; i < len(array); i++ {
		if cmp.Less(array[i], array[i-1]) {
			return false
		}
	}
	return true
}

func (array Dynamic[T]) LinearSearch(value T) int {
	return searching.LinearSearch(array, value)
}

func (array Dynamic[T]) BinarySearch(value T) int {
	if !array.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(array, value)
}
