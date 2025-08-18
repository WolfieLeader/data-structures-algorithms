package arraystack

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
)

type ArrayStack[T cmp.Ordered] struct {
	data dynamic.Dynamic[T]
}

func New[T cmp.Ordered]() *ArrayStack[T] {
	return &ArrayStack[T]{data: dynamic.New[T]()}
}
