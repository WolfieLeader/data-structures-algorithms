package bst

import "cmp"

type iterativeBST[T cmp.Ordered] interface {
	ContainsI(value T) bool
	InsertI(values ...T)
	DeleteI(value T) bool

	TraverseInOrderI(fn func(value T))
	TraversePreOrderI(fn func(value T))
	TraversePostOrderI(fn func(value T))
	TraverseBreadthFirst(fn func(value T))

	MinI() (T, bool)
	MaxI() (T, bool)
	HeightI() int
	PredecessorI(value T) (T, bool)
	SuccessorI(value T) (T, bool)
	ToSliceI() []T
	EqualI(other *BST[T]) bool
	CopyI() *BST[T]
}

type recursiveBST[T cmp.Ordered] interface {
	ContainsR(value T) bool
	InsertR(values ...T)
	DeleteR(value T) bool

	TraverseInOrderR(fn func(value T))
	TraversePreOrderR(fn func(value T))
	TraversePostOrderR(fn func(value T))

	MinR() (T, bool)
	MaxR() (T, bool)
	HeightR() int
	PredecessorR(value T) (T, bool)
	SuccessorR(value T) (T, bool)
	IsBalanced() bool
	ToSliceR() []T
	EqualR(other *BST[T]) bool
	CopyR() *BST[T]
	String() string
}

type BinarySearchTree[T cmp.Ordered] interface {
	iterativeBST[T]
	recursiveBST[T]
	Size() int
	IsEmpty() bool
	Root() (T, bool)
	Clear()
}

var _ BinarySearchTree[int] = (*BST[int])(nil)
