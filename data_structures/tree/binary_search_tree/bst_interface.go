package bst

type BinarySearchTree[T comparable] interface {
	Insert(value T) bool
	Delete(value T) bool
	Search(value T) bool
	Contains(value T) bool

	TraverseInOrder(fn func(value T))
	TraversePreOrder(fn func(value T))
	TraversePostOrder(fn func(value T))

	Size() int
	IsEmpty() bool
	Root() T
	Height() int
	IsBalanced() bool
	Min() T
	Max() T
	Clear()
	Copy() *BST[T]
}
