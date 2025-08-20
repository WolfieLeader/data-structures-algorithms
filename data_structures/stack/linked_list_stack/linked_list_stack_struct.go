package liststack

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/singly"
)

type LinkedListStack[T cmp.Ordered] struct {
	data *singly.Singly[T]
}

func New[T cmp.Ordered]() *LinkedListStack[T] {
	return &LinkedListStack[T]{data: singly.New[T]()}
}
