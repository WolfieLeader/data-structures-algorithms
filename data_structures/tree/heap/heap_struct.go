package heap

import "cmp"

type Heap[T any] struct {
	data      []T
	less      func(a, b T) bool
	isMinHeap bool
}

func New[T cmp.Ordered](isMinHeap bool) *Heap[T] {
	data := make([]T, 0)
	lessFn := func(a, b T) bool { return a < b }
	if !isMinHeap {
		lessFn = func(a, b T) bool { return a > b }
	}
	return &Heap[T]{
		data:      data,
		less:      lessFn,
		isMinHeap: isMinHeap,
	}
}

func NewWithFunc[T any](less func(a, b T) bool, isMinHeap bool) *Heap[T] {
	return &Heap[T]{
		data:      make([]T, 0),
		less:      less,
		isMinHeap: isMinHeap,
	}
}
