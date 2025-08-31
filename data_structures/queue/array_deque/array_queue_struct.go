package arraydeque

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
)

type ArrayDeque[T cmp.Ordered] struct {
	data *dynamic.Dynamic[T]
}

func New[T cmp.Ordered]() *ArrayDeque[T] {
	return &ArrayDeque[T]{data: dynamic.New[T]()}
}
