package liststack

import "cmp"

type LinkedListStack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ListStack[T]
}

var _ LinkedListStack[int] = (*ListStack[int])(nil)
