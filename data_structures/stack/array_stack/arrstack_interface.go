package arraystack

import "cmp"

type ArrayStack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Copy() *ArrStack[T]
}
