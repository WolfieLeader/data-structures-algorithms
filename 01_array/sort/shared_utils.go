package sort

import "cmp"

type array[T cmp.Ordered] []T

func copyArray[T cmp.Ordered](arr []T) (array[T], int) {
	length := len(arr)
	out := make(array[T], length)
	copy(out, arr)
	return out, length
}

func (array *array[T]) swap(i, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
