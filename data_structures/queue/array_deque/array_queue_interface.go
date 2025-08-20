package arraydeque

import "cmp"

type Deque[T cmp.Ordered] interface {
	EnqueueFirst(values ...T)
	EnqueueLast(values ...T)
	DequeueFirst() (T, bool)
	DequeueLast() (T, bool)
	PeekFirst() (T, bool)
	PeekLast() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ArrayDeque[T]
}

var _ Deque[int] = (*ArrayDeque[int])(nil)
