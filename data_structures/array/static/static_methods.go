package static

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (arr Static[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(arr) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}

	return arr[i], nil
}

func (arr *Static[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*arr) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	(*arr)[i] = value
	return nil
}

func (arr *Static[T]) Replace(values ...T) {
	var zero T
	arr.Fill(zero)

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy((*arr)[:], values)
}

func (arr *Static[T]) Fill(value T) {
	for i := range *arr {
		(*arr)[i] = value
	}
}

func (arr *Static[T]) Clear() {
	*arr = Static[T]{}
}

func (arr Static[T]) Length() int {
	return len(arr)
}

func (arr Static[T]) IsSorted() bool {
	for i := 1; i < len(arr); i++ {
		if utils.Is(arr[i-1], utils.GreaterThan, arr[i]) {
			return false
		}
	}
	return true
}

func (arr Static[T]) Search(value T) int {
	return searching.LinearSearch(arr[:], value)
}

func (arr Static[T]) BinarySearch(value T) int {
	if !arr.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(arr[:], value)
}

func (arr Static[T]) Contains(value T) bool {
	return arr.Search(value) != -1
}

func (arr Static[T]) Traverse(fn func(i int, value T) bool) {
	for index, value := range arr {
		if !fn(index, value) {
			break
		}
	}
}

func (arr *Static[T]) Swap(i, j int) error {
	length := len(*arr)
	if i < 0 || i >= length || j < 0 || j >= length {
		return fmt.Errorf("index out of bounds: %d or %d", i, j)
	}

	if i == j {
		return nil
	}

	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	return nil
}

func (arr Static[T]) ToSlice() []T {
	out := make([]T, len(arr))
	copy(out, arr[:])
	return out
}

func (arr Static[T]) Copy() Static[T] {
	var out Static[T]
	copy(out[:], arr[:])
	return out
}

func (arr Static[T]) Reverse() Static[T] {
	out := arr.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
