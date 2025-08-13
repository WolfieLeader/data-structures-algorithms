package sort

import "cmp"

type Ordered = cmp.Ordered
type array[T Ordered] []T

func less[T Ordered](a, b T) bool {
	return cmp.Less(a, b)
}

func copyArray[T Ordered](arr []T) (array[T], int) {
	length := len(arr)
	out := make(array[T], length)
	copy(out, arr)
	return out, length
}

func (array *array[T]) swap(i, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
