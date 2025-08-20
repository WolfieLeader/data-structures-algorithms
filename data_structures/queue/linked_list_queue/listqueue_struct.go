package listqueue

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/singly"
)

type ListQueue[T cmp.Ordered] struct {
	data *singly.Singly[T]
}

func New[T cmp.Ordered]() *ListQueue[T] {
	return &ListQueue[T]{data: singly.New[T]()}
}
