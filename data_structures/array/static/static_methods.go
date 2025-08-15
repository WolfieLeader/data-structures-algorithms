package static

import (
	"cmp"
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
)

func (a Static[T]) Get(i int) (T, error) {
	if i < 0 || i >= len(a) {
		var zeroValue T
		return zeroValue, fmt.Errorf("index out of bounds: %d", i)
	}

	return a[i], nil
}

func (a *Static[T]) Set(i int, value T) error {
	if i < 0 || i >= len(*a) {
		return fmt.Errorf("index out of bounds: %d", i)
	}

	(*a)[i] = value
	return nil
}

func (a *Static[T]) Replace(values ...T) {
	var zeroValue T
	a.Fill(zeroValue)

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy((*a)[:], values)
}

func (a *Static[T]) Fill(value T) {
	for i := range *a {
		(*a)[i] = value
	}
}

func (a *Static[T]) Clear() {
	*a = Static[T]{}
}

func (a Static[T]) Length() int {
	return len(a)
}

func (a Static[T]) IsSorted() bool {
	for i := 1; i < len(a); i++ {
		if cmp.Less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}

func (a Static[T]) LinearSearch(value T) int {
	return searching.LinearSearch(a[:], value)
}

func (a Static[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a[:], value)
}

func (a Static[T]) Contains(value T) bool {
	return a.LinearSearch(value) != -1
}

func (a Static[T]) Traverse(fn func(i int, value T) bool) {
	for index, value := range a {
		if !fn(index, value) {
			break
		}
	}
}

func (a *Static[T]) Swap(i, j int) error {
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

func (a Static[T]) ToSlice() []T {
	out := make([]T, len(a))
	copy(out, a[:])
	return out
}

func (a Static[T]) Copy() Static[T] {
	var out Static[T]
	copy(out[:], a[:])
	return out
}

func (a Static[T]) Reverse() Static[T] {
	out := a.Copy()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return out
}
