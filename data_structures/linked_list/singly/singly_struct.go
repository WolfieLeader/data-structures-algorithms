package singly

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	next  *Node[T]
}

type Singly[T cmp.Ordered] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func New[T cmp.Ordered]() *Singly[T] {
	return &Singly[T]{}
}
