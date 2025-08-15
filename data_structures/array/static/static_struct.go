package static

import "cmp"

const SIZE = 5

type Static[T cmp.Ordered] [SIZE]T

func New[T cmp.Ordered](values ...T) Static[T] {
	var staticArray Static[T]
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(staticArray[:], values)
	return staticArray
}
