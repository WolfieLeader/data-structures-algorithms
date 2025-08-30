package bst

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	Contains(value T) bool
	Insert(value T) bool
	Delete(value T) bool
	RecursiveContains(value T) bool
	RecursiveInsert(value T) bool
	RecursiveDelete(value T) bool

	TraverseDFSInOrder(fn func(value T))
	TraverseDFSPreOrder(fn func(value T))
	TraverseDFSPostOrder(fn func(value T))
	IterativeTraverseDFSInOrder(fn func(value T))
	IterativeTraverseDFSPreOrder(fn func(value T))
	IterativeTraverseDFSPostOrder(fn func(value T))

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
