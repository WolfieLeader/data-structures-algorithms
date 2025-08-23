package arraystack

import "cmp"

type Stack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ArrayStack[T]
}

var _ Stack[int] = (*ArrayStack[int])(nil)
