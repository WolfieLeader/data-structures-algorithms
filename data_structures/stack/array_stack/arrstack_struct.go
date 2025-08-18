package arraystack

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
)

type ArrStack[T cmp.Ordered] struct {
	data dynamic.Dynamic[T]
}

func New[T cmp.Ordered]() *ArrStack[T] {
	return &ArrStack[T]{data: dynamic.New[T]()}
}
