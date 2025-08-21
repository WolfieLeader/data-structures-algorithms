package static

import "cmp"

const SIZE = 5

type Static[T cmp.Ordered] struct {
	data [SIZE]T
}

func New[T cmp.Ordered](values ...T) Static[T] {
	var data [SIZE]T

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(data[:], values)
	return Static[T]{data: data}
}
