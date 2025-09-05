package avl

import "cmp"

type avlTree[T cmp.Ordered] interface {
	Contains(value T) bool
	Insert(values ...T)
	Delete(value T) bool
	Size() int
	IsEmpty() bool
	Clear()
	Min() (T, bool)
	Max() (T, bool)
	Height() int
	BalanceFactor() int
	ToSlice() []T
	Equal(other *AVLTree[T]) bool
	Copy() *AVLTree[T]
	String() string
}

var _ avlTree[int] = (*AVLTree[int])(nil)
