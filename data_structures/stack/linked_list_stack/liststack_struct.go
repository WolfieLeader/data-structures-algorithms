package liststack

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/singly"
)

type ListStack[T cmp.Ordered] struct {
	data *singly.Singly[T]
}

func New[T cmp.Ordered]() *ListStack[T] {
	return &ListStack[T]{data: singly.New[T]()}
}
