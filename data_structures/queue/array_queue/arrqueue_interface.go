package arrqueue

import "cmp"

type ArrayQueue[T cmp.Ordered] interface {
	Enqueue(values ...T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ArrQueue[T]
}
