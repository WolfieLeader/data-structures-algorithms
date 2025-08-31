package dynamic

import "cmp"

type Dynamic[T cmp.Ordered] struct {
	data []T
}

func New[T cmp.Ordered](values ...T) *Dynamic[T] {
	return &Dynamic[T]{data: values}
}
