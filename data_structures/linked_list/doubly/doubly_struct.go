package doubly

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}

type doublyLinkedList[T cmp.Ordered] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func New[T cmp.Ordered]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{}
}
