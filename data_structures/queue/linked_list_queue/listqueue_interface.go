package listqueue

import "cmp"

type LinkedListQueue[T cmp.Ordered] interface {
	Enqueue(values ...T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ListQueue[T]
}
