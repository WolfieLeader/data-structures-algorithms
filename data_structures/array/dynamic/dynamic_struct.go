package dynamic

import "cmp"

type Dynamic[T cmp.Ordered] []T

func New[T cmp.Ordered](values ...T) Dynamic[T] {
	return append(Dynamic[T]{}, values...)
}
