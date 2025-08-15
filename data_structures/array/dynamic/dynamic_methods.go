package dynamic

import (
	"cmp"
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
)

func (a dynamicArray[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(a) {
		var zeroValue T
		return zeroValue, fmt.Errorf("index out of bounds: %d", i)
	}
	return a[i], nil
}

func (a *dynamicArray[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*a) {
		return fmt.Errorf("index out of bounds: %d", i)
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*a)[i] = value
	return nil
}

func (a *dynamicArray[T]) Replace(values ...T) {
	*a = append((*a)[:0], values...)
}

func (a *dynamicArray[T]) Append(values ...T) {
	*a = append(*a, values...)
}

func (a *dynamicArray[T]) Prepend(values ...T) {
	*a = append(values, *a...)
}

func (a *dynamicArray[T]) Insert(i int, value T) error {
	if i < 0 || i > len(*a) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	var zeroValue T
	*a = append(*a, zeroValue)
	copy((*a)[i+1:], (*a)[i:])
	(*a)[i] = value

	return nil
}

func (a *dynamicArray[T]) Delete(i int) (T, error) {
	if i < 0 || i >= len(*a) {
		var zeroValue T
		return zeroValue, fmt.Errorf("index out of bounds: %d", i)
	}

	value := (*a)[i]
	*a = append((*a)[:i], (*a)[i+1:]...)
	return value, nil
}

func (a *dynamicArray[T]) Fill(value T) {
	for i := range *a {
		(*a)[i] = value
	}
}

func (a *dynamicArray[T]) Clear() {
	*a = dynamicArray[T]{}
}

func (a dynamicArray[T]) Length() int {
	return len(a)
}

func (a dynamicArray[T]) Capacity() int {
	return cap(a)
}

func (a dynamicArray[T]) IsSorted() bool {
	for i := 1; i < len(a); i++ {
		if cmp.Less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}

func (a dynamicArray[T]) LinearSearch(value T) int {
	return searching.LinearSearch(a, value)
}

func (a dynamicArray[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a, value)
}

func (a dynamicArray[T]) Contains(value T) bool {
	return searching.LinearSearch(a, value) != -1
}

func (a dynamicArray[T]) Traverse(fn func(i int, value T) bool) {
	for i, value := range a {
		if !fn(i, value) {
			break
		}
	}
}

func (a *dynamicArray[T]) Swap(i, j int) error {
	length := len(*a)
	if i < 0 || i >= length || j < 0 || j >= length {
		return fmt.Errorf("index out of bounds: %d, %d", i, j)
	}

	if i == j {
		return nil
	}

	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	return nil
}

func (a dynamicArray[T]) Slice(start, end int) (dynamicArray[T], error) {
	if start < 0 || end > len(a) || start >= end {
		return nil, fmt.Errorf("invalid slice range: %d to %d", start, end)
	}

	out := make(dynamicArray[T], end-start)
	copy(out, a[start:end])
	return out, nil
}

func (a dynamicArray[T]) Copy() dynamicArray[T] {
	out := make(dynamicArray[T], len(a))
	copy(out, a)
	return out
}

func (a dynamicArray[T]) Reverse() dynamicArray[T] {
	out := a.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
