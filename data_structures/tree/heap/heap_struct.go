package heap

import "cmp"

type HeapType int

const (
	minHeap HeapType = iota
	maxHeap
)

type Heap[T any] struct {
	data     []T
	less     func(a, b T) bool
	heapType HeapType
}

func New[T cmp.Ordered](isMinHeap bool) *Heap[T] {
	data := make([]T, 0)
	if isMinHeap {
		return &Heap[T]{
			data:     data,
			less:     func(a, b T) bool { return a < b },
			heapType: minHeap,
		}
	}
	return &Heap[T]{
		data:     data,
		less:     func(a, b T) bool { return a > b },
		heapType: maxHeap,
	}
}

func NewWithFunc[T any](less func(a, b T) bool, isMinHeap bool) *Heap[T] {
	if isMinHeap {
		return &Heap[T]{
			data:     make([]T, 0),
			less:     less,
			heapType: minHeap,
		}
	}
	return &Heap[T]{
		data:     make([]T, 0),
		less:     less,
		heapType: maxHeap,
	}
}
