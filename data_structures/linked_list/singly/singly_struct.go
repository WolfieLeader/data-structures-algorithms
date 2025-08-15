package singly

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Singly[T comparable] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func New[T comparable]() *Singly[T] {
	return &Singly[T]{}
}
