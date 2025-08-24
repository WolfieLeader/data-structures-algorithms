package liststack

import "cmp"

type Stack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *LinkedListStack[T]
	String() string
}

var _ Stack[int] = (*LinkedListStack[int])(nil)
