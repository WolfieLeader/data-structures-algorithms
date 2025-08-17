package dynamic

import "cmp"

type dynamic[T cmp.Ordered] []T

func New[T cmp.Ordered](values ...T) dynamic[T] {
	return append(dynamic[T]{}, values...)
}
