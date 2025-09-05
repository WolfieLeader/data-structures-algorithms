package listqueue

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/singly"
)

type LinkedListQueue[T cmp.Ordered] struct {
	data *singly.SinglyLinkedList[T]
}

func New[T cmp.Ordered]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{data: singly.New[T]()}
}
