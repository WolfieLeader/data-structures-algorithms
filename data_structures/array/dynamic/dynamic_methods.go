package dynamic

import (
	"fmt"
	"slices"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
)

func (a *Dynamic[T]) Length() int   { return len(a.data) }
func (a *Dynamic[T]) Capacity() int { return cap(a.data) }
func (a *Dynamic[T]) IsEmpty() bool { return len(a.data) == 0 }
func (a *Dynamic[T]) Clear()        { *a = Dynamic[T]{} }

func (a *Dynamic[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, false
	}
	return a.data[index], true
}

func (a *Dynamic[T]) Set(index int, value T) bool {
	if index < 0 || index >= len(a.data) {
		return false
	}
	a.data[index] = value
	return true
}

func (a *Dynamic[T]) Replace(values ...T) {
	a.data = append(a.data[:0], values...)
}

func (a *Dynamic[T]) Append(values ...T) {
	a.data = append(a.data, values...)
}

func (a *Dynamic[T]) Prepend(values ...T) {
	if len(values) == 0 {
		return
	}
	out := make([]T, 0, len(values)+len(a.data))
	out = append(out, values...)
	a.data = append(out, a.data...)
}

func (a *Dynamic[T]) Insert(index int, value T) bool {
	if index < 0 || index > len(a.data) {
		return false
	}

	var zero T
	a.data = append(a.data, zero)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value

	return true
}

func (a *Dynamic[T]) Delete(index int) (T, bool) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, false
	}

	value := a.data[index]
	a.data = append(a.data[:index], a.data[index+1:]...)
	return value, true
}

func (a *Dynamic[T]) Fill(value T) {
	for i := range a.data {
		a.data[i] = value
	}
}

func (a *Dynamic[T]) IsSorted() bool {
	for i := 1; i < len(a.data); i++ {
		if a.data[i-1] > a.data[i] {
			return false
		}
	}
	return true
}

func (a *Dynamic[T]) Search(value T) int {
	return searching.LinearSearch(a.data, value)
}

func (a *Dynamic[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a.data, value)
}

func (a *Dynamic[T]) Contains(value T) bool {
	return a.Search(value) != -1
}

func (a *Dynamic[T]) Traverse(fn func(index int, value T)) {
	for i, value := range a.data {
		fn(i, value)
	}
}

func (a *Dynamic[T]) Swap(index1, index2 int) bool {
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

func (a *Dynamic[T]) Copy() *Dynamic[T] {
	return &Dynamic[T]{data: a.ToSlice()}
}

func (a *Dynamic[T]) Equal(other *Dynamic[T]) bool {
	if a == other {
		return true
	}
	if a == nil || other == nil {
		return false
	}
	return slices.Equal(a.data, other.data)
}

func (a *Dynamic[T]) Between(start, end int) *Dynamic[T] {
	if start < 0 || end > len(a.data) || start > end {
		return &Dynamic[T]{}
	}

	out := make([]T, end-start)
	copy(out, a.data[start:end])
	return &Dynamic[T]{data: out}
}

func (a *Dynamic[T]) Reverse() *Dynamic[T] {
	out := a.ToSlice()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return &Dynamic[T]{data: out}
}

func (a *Dynamic[T]) ToSlice() []T {
	out := make([]T, len(a.data))
	copy(out, a.data)
	return out
}

func (a *Dynamic[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}
