package static

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (a staticArray[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(a) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", i)
	}

	return a[i], nil
}

func (a *staticArray[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*a) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	(*a)[i] = value
	return nil
}

func (a *staticArray[T]) Replace(values ...T) {
	var zero T
	a.Fill(zero)

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy((*a)[:], values)
}

func (a *staticArray[T]) Fill(value T) {
	for i := range *a {
		(*a)[i] = value
	}
}

func (a *staticArray[T]) Clear() {
	*a = staticArray[T]{}
}

func (a staticArray[T]) Length() int {
	return len(a)
}

func (a staticArray[T]) IsSorted() bool {
	for i := 1; i < len(a); i++ {
		if utils.Is(a[i-1], utils.GreaterThan, a[i]) {
			return false
		}
	}
	return true
}

func (a staticArray[T]) LinearSearch(value T) int {
	return searching.LinearSearch(a[:], value)
}

func (a staticArray[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a[:], value)
}

func (a staticArray[T]) Contains(value T) bool {
	return a.LinearSearch(value) != -1
}

func (a staticArray[T]) Traverse(fn func(i int, value T) bool) {
	for index, value := range a {
		if !fn(index, value) {
			break
		}
	}
}

func (a *staticArray[T]) Swap(i, j int) error {
	length := len(*a)
	if i < 0 || i >= length || j < 0 || j >= length {
		return fmt.Errorf("index out of bounds: %d or %d", i, j)
	}

	if i == j {
		return nil
	}

	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	return nil
}

func (a staticArray[T]) ToSlice() []T {
	out := make([]T, len(a))
	copy(out, a[:])
	return out
}

func (a staticArray[T]) Copy() staticArray[T] {
	var out staticArray[T]
	copy(out[:], a[:])
	return out
}

func (a staticArray[T]) Reverse() staticArray[T] {
	out := a.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
