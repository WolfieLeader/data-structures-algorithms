package dynamic

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/searching_algorithms/searching"
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (a Dynamic[T]) Get(i int) (T, bool) {
	if i < 0 || i >= len(a.data) {
		var zero T
		return zero, false
	}
	return a.data[i], true
}

func (a *Dynamic[T]) Set(i int, value T) bool {
	if i < 0 || i >= len(a.data) {
		return false
	}
	a.data[i] = value
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

func (a *Dynamic[T]) Insert(i int, value T) bool {
	if i < 0 || i > len(a.data) {
		return false
	}

	var zero T
	a.data = append(a.data, zero)
	copy(a.data[i+1:], a.data[i:])
	a.data[i] = value

	return true
}

func (a *Dynamic[T]) Delete(i int) (T, bool) {
	if i < 0 || i >= len(a.data) {
		var zero T
		return zero, false
	}

	value := a.data[i]
	a.data = append(a.data[:i], a.data[i+1:]...)
	return value, true
}

func (a *Dynamic[T]) Fill(value T) {
	for i := range a.data {
		a.data[i] = value
	}
}

func (a *Dynamic[T]) Clear() {
	*a = Dynamic[T]{}
}

func (a Dynamic[T]) Length() int {
	return len(a.data)
}

func (a Dynamic[T]) Capacity() int {
	return cap(a.data)
}

func (a Dynamic[T]) IsSorted() bool {
	for i := 1; i < len(a.data); i++ {
		if utils.Is(a.data[i-1], utils.GreaterThan, a.data[i]) {
			return false
		}
	}
	return true
}

func (a Dynamic[T]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a Dynamic[T]) Search(value T) int {
	return searching.LinearSearch(a.data, value)
}

func (a Dynamic[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a.data, value)
}

func (a Dynamic[T]) Contains(value T) bool {
	return a.Search(value) != -1
}

func (a Dynamic[T]) Traverse(fn func(i int, value T) bool) {
	for i, value := range a.data {
		if !fn(i, value) {
			break
		}
	}
}

func (a *Dynamic[T]) Swap(i, j int) bool {
	length := len(a.data)
	if i < 0 || i >= length || j < 0 || j >= length {
		return false
	}

	if i == j {
		return true
	}

	a.data[i], a.data[j] = a.data[j], a.data[i]
	return true
}

func (a Dynamic[T]) Copy() Dynamic[T] {
	return Dynamic[T]{data: a.ToSlice()}
}

func (a Dynamic[T]) Between(start, end int) (Dynamic[T], bool) {
	if start < 0 || end > len(a.data) || start > end {
		return Dynamic[T]{}, false
	}

	out := make([]T, end-start)
	copy(out, a.data[start:end])
	return Dynamic[T]{data: out}, true
}

func (a Dynamic[T]) Reverse() Dynamic[T] {
	out := a.ToSlice()
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return Dynamic[T]{data: out}
}

func (a Dynamic[T]) ToSlice() []T {
	out := make([]T, len(a.data))
	copy(out, a.data)
	return out
}

func (a Dynamic[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}
