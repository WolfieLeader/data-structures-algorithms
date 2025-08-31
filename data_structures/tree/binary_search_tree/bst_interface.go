package bst

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	ContainsI(value T) bool
	InsertI(value T) bool
	DeleteI(value T) bool
	ContainsR(value T) bool
	InsertR(value T) bool
	DeleteR(value T) bool

	TraverseInOrderI(fn func(value T))
	TraversePreOrderI(fn func(value T))
	TraversePostOrderI(fn func(value T))
	TraverseBreadthFirstI(fn func(value T))
	TraverseInOrderR(fn func(value T))
	TraversePreOrderR(fn func(value T))
	TraversePostOrderR(fn func(value T))

	Size() int
	IsEmpty() bool
	Root() (T, bool)
	MinI() (T, bool)
	MaxI() (T, bool)
	MinR() (T, bool)
	MaxR() (T, bool)
	Height() int
	IsBalanced() bool
	Clear()
	Copy() *BST[T]
}

var _ BinarySearchTree[int] = (*BST[int])(nil)
