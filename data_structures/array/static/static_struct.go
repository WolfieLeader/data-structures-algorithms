package static

import "cmp"

const SIZE = 5

type staticArray[T cmp.Ordered] [SIZE]T

func New[T cmp.Ordered](values ...T) staticArray[T] {
	var staticArray staticArray[T]

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(staticArray[:], values)
	return staticArray
}
