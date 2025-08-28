package bst

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	Contains(value T) bool
	RecursiveContains(value T) bool
	Insert(value T) bool
	RecursiveInsert(value T) bool
	Delete(value T) bool
	RecursiveDelete(value T) bool

	TraverseInOrder(fn func(value T))
	TraversePreOrder(fn func(value T))
	TraversePostOrder(fn func(value T))

	Size() int
	IsEmpty() bool
	Root() (T, bool)
	Min() (T, bool)
	Max() (T, bool)
	Height() int
	IsBalanced() bool
	Clear()
	Copy() *BST[T]
}

var _ BinarySearchTree[int] = (*BST[int])(nil)
