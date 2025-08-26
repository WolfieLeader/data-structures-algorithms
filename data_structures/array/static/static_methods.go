package static

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
)

func (a Static[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, false
	}

	return a.data[index], true
}

func (a *Static[T]) Set(index int, value T) bool {
	if index < 0 || index >= len(a.data) {
		return false
	}

	a.data[index] = value
	return true
}

func (a *Static[T]) Replace(values ...T) {
	var zero T
	a.Fill(zero)

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(a.data[:], values)
}

func (a *Static[T]) Fill(value T) {
	for i := range a.data {
		a.data[i] = value
	}
}

func (a *Static[T]) Clear() {
	*a = Static[T]{}
}

func (a Static[T]) Length() int {
	return len(a.data)
}

func (a Static[T]) IsSorted() bool {
	for i := 1; i < len(a.data); i++ {
		if a.data[i-1] > a.data[i] {
			return false
		}
	}
	return true
}

func (a Static[T]) Search(value T) int {
	return searching.LinearSearch(a.data[:], value)
}

func (a Static[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a.data[:], value)
}

func (a Static[T]) Contains(value T) bool {
	return a.Search(value) != -1
}

func (a Static[T]) Traverse(fn func(index int, value T)) {
	for index, value := range a.data {
		fn(index, value)
	}
}

func (a *Static[T]) Swap(index1, index2 int) bool {
	length := len(a.data)
	if index1 < 0 || index1 >= length || index2 < 0 || index2 >= length {
		return false
	}

	if index1 == index2 {
		return true
	}

	a.data[index1], a.data[index2] = a.data[index2], a.data[index1]
	return true
}

func (a Static[T]) ToSlice() []T {
	out := make([]T, len(a.data))
	copy(out, a.data[:])
	return out
}

func (a Static[T]) Copy() Static[T] {
	var out Static[T]
	copy(out.data[:], a.data[:])
	return out
}

func (a Static[T]) Reverse() Static[T] {
	out := a.data
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return Static[T]{data: out}
}

func (a Static[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}
