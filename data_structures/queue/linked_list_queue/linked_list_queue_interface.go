package listqueue

import "cmp"

type Queue[T cmp.Ordered] interface {
	Enqueue(values ...T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *LinkedListQueue[T]
}

var _ Queue[int] = (*LinkedListQueue[int])(nil)
