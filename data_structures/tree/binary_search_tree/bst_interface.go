package bst

import "cmp"

type iterativeBST[T cmp.Ordered] interface {
	ContainsI(value T) bool
	InsertI(value T) bool
	DeleteI(value T) bool

	TraverseInOrderI(fn func(value T))
	TraversePreOrderI(fn func(value T))
	TraversePostOrderI(fn func(value T))
	TraverseBreadthFirst(fn func(value T))

	MinI() (T, bool)
	MaxI() (T, bool)
	HeightI() int
	CopyI() *BST[T]
}

type recursiveBST[T cmp.Ordered] interface {
	ContainsR(value T) bool
	InsertR(value T) bool
	DeleteR(value T) bool

	TraverseInOrderR(fn func(value T))
	TraversePreOrderR(fn func(value T))
	TraversePostOrderR(fn func(value T))

	MinR() (T, bool)
	MaxR() (T, bool)
	HeightR() int
	IsBalanced() bool
	IsSymmetric() bool
	CopyR() *BST[T]
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
