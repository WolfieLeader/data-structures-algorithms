package arrqueue

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
)

type ArrQueue[T cmp.Ordered] struct {
	data dynamic.Dynamic[T]
}

func New[T cmp.Ordered]() *ArrQueue[T] {
	return &ArrQueue[T]{data: dynamic.New[T]()}
}
