package static

import "cmp"

const SIZE = 5

type static[T cmp.Ordered] [SIZE]T

func New[T cmp.Ordered](values ...T) static[T] {
	var staticArray static[T]

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(staticArray[:], values)
	return staticArray
}
