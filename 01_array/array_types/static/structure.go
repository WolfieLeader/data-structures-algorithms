package static

import "cmp"

const SIZE = 5

type StaticArrayType[T cmp.Ordered] [SIZE]T

func New[T cmp.Ordered](values ...T) StaticArrayType[T] {
	var staticArray StaticArrayType[T]
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(staticArray[:], values)
	return staticArray
}
