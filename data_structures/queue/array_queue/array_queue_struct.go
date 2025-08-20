package arrayqueue

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
)

type ArrayQueue[T cmp.Ordered] struct {
	data dynamic.Dynamic[T]
}

func New[T cmp.Ordered]() *ArrayQueue[T] {
	return &ArrayQueue[T]{data: dynamic.New[T]()}
}
