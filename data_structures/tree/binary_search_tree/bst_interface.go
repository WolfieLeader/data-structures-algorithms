package bst

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	Insert(value T) bool
	Delete(value T) bool
	Contains(value T) bool

	TraverseInOrder(fn func(value T))
	TraversePreOrder(fn func(value T))
	TraversePostOrder(fn func(value T))

	Size() int
	IsEmpty() bool
	Root() T
	Min() T
	Max() T
	Height() int
	IsBalanced() bool
	Clear()
	Copy() *BST[T]
}

var _ BinarySearchTree[int] = (*BST[int])(nil)
