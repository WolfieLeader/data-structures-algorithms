package bst

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	left  *Node[T]
	right *Node[T]
}

type BST[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func New[T cmp.Ordered]() *BST[T] {
	return &BST[T]{}
}
