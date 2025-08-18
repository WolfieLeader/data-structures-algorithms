package dynamic

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (arr Dynamic[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(arr) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}
	return arr[i], nil
}

func (arr *Dynamic[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*arr) {
		return fmt.Errorf("index out of bounds: %d", i)
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*arr)[i] = value
	return nil
}

func (arr *Dynamic[T]) Replace(values ...T) {
	*arr = append((*arr)[:0], values...)
}

func (arr *Dynamic[T]) Append(values ...T) {
	*arr = append(*arr, values...)
}

func (arr *Dynamic[T]) Prepend(values ...T) {
	*arr = append(values, *arr...)
}

func (arr *Dynamic[T]) Insert(i int, value T) error {
	if i < 0 || i > len(*arr) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	var zero T
	*arr = append(*arr, zero)
	copy((*arr)[i+1:], (*arr)[i:])
	(*arr)[i] = value

	return nil
}

func (arr *Dynamic[T]) Delete(i int) (T, error) {
	if i < 0 || i >= len(*arr) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}

	value := (*arr)[i]
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
	return value, nil
}

func (arr *Dynamic[T]) Fill(value T) {
	for i := range *arr {
		(*arr)[i] = value
	}
}

func (arr *Dynamic[T]) Clear() {
	*arr = Dynamic[T]{}
}

func (arr Dynamic[T]) Length() int {
	return len(arr)
}

func (arr Dynamic[T]) Capacity() int {
	return cap(arr)
}

func (arr Dynamic[T]) IsSorted() bool {
	for i := 1; i < len(arr); i++ {
		if utils.Is(arr[i-1], utils.GreaterThan, arr[i]) {
			return false
		}
	}
	return true
}

func (arr Dynamic[T]) Search(value T) int {
	return searching.LinearSearch(arr, value)
}

func (arr Dynamic[T]) BinarySearch(value T) int {
	if !arr.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(arr, value)
}

func (arr Dynamic[T]) Contains(value T) bool {
	return arr.Search(value) != -1
}

func (arr Dynamic[T]) Traverse(fn func(i int, value T) bool) {
	for i, value := range arr {
		if !fn(i, value) {
			break
		}
	}
}

func (arr *Dynamic[T]) Swap(i, j int) error {
	length := len(*arr)
	if i < 0 || i >= length || j < 0 || j >= length {
		return fmt.Errorf("index out of bounds: %d, %d", i, j)
	}

	if i == j {
		return nil
	}

	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	return nil
}

func (arr Dynamic[T]) Slice(start, end int) (Dynamic[T], error) {
	if start < 0 || end > len(arr) || start >= end {
		return nil, fmt.Errorf("invalid slice range: %d to %d", start, end)
	}

	out := make(Dynamic[T], end-start)
	copy(out, arr[start:end])
	return out, nil
}

func (arr Dynamic[T]) Copy() Dynamic[T] {
	out := make(Dynamic[T], len(arr))
	copy(out, arr)
	return out
}

func (arr Dynamic[T]) Reverse() Dynamic[T] {
	out := arr.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
