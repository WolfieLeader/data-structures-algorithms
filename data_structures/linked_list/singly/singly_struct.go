package singly

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	next  *Node[T]
}

type singly[T cmp.Ordered] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func newNode[T cmp.Ordered](value T, next *Node[T]) *Node[T] {
	return &Node[T]{Value: value, next: next}
}

func New[T cmp.Ordered]() *singly[T] {
	return &singly[T]{}
}
