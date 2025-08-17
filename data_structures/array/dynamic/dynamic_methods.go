package dynamic

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (arr dynamic[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(arr) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}
	return arr[i], nil
}

func (arr *dynamic[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*arr) {
		return fmt.Errorf("index out of bounds: %d", i)
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*arr)[i] = value
	return nil
}

func (arr *dynamic[T]) Replace(values ...T) {
	*arr = append((*arr)[:0], values...)
}

func (arr *dynamic[T]) Append(values ...T) {
	*arr = append(*arr, values...)
}

func (arr *dynamic[T]) Prepend(values ...T) {
	*arr = append(values, *arr...)
}

func (arr *dynamic[T]) Insert(i int, value T) error {
	if i < 0 || i > len(*arr) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	var zero T
	*arr = append(*arr, zero)
	copy((*arr)[i+1:], (*arr)[i:])
	(*arr)[i] = value

	return nil
}

func (arr *dynamic[T]) Delete(i int) (T, error) {
	if i < 0 || i >= len(*arr) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}

	value := (*arr)[i]
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
	return value, nil
}

func (arr *dynamic[T]) Fill(value T) {
	for i := range *arr {
		(*arr)[i] = value
	}
}

func (arr *dynamic[T]) Clear() {
	*arr = dynamic[T]{}
}

func (arr dynamic[T]) Length() int {
	return len(arr)
}

func (arr dynamic[T]) Capacity() int {
	return cap(arr)
}

func (arr dynamic[T]) IsSorted() bool {
	for i := 1; i < len(arr); i++ {
		if utils.Is(arr[i-1], utils.GreaterThan, arr[i]) {
			return false
		}
	}
	return true
}

func (arr dynamic[T]) Search(value T) int {
	return searching.LinearSearch(arr, value)
}

func (arr dynamic[T]) BinarySearch(value T) int {
	if !arr.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(arr, value)
}

func (arr dynamic[T]) Contains(value T) bool {
	return arr.Search(value) != -1
}

func (arr dynamic[T]) Traverse(fn func(i int, value T) bool) {
	for i, value := range arr {
		if !fn(i, value) {
			break
		}
	}
}

func (arr *dynamic[T]) Swap(i, j int) error {
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

func (arr dynamic[T]) Slice(start, end int) (dynamic[T], error) {
	if start < 0 || end > len(arr) || start >= end {
		return nil, fmt.Errorf("invalid slice range: %d to %d", start, end)
	}

	out := make(dynamic[T], end-start)
	copy(out, arr[start:end])
	return out, nil
}

func (arr dynamic[T]) Copy() dynamic[T] {
	out := make(dynamic[T], len(arr))
	copy(out, arr)
	return out
}

func (arr dynamic[T]) Reverse() dynamic[T] {
	out := arr.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
