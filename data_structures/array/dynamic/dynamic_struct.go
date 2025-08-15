package dynamic

import "cmp"

type dynamicArray[T cmp.Ordered] []T

func New[T cmp.Ordered](values ...T) dynamicArray[T] {
	return append(dynamicArray[T]{}, values...)
}
