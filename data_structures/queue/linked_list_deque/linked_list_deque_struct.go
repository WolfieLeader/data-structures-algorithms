package listdeque

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/doubly"
)

type LinkedListDeque[T cmp.Ordered] struct {
	data *doubly.Doubly[T]
}

func New[T cmp.Ordered]() *LinkedListDeque[T] {
	return &LinkedListDeque[T]{data: doubly.New[T]()}
}
