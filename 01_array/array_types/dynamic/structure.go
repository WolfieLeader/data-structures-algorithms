package dynamic

import "cmp"

type DynamicArrayType[T cmp.Ordered] []T

func New[T cmp.Ordered](values ...T) DynamicArrayType[T] {
	return append(DynamicArrayType[T]{}, values...)
}
