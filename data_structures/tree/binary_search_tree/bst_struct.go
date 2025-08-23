package bst

type Node[T comparable] struct {
	Value T
	left  *Node[T]
	right *Node[T]
}

type BST[T comparable] struct {
	root *Node[T]
	size int
}

func New[T comparable]() *BST[T] {
	return &BST[T]{}
}
