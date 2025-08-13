package sort

import "cmp"

func Bubble[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for index := 0; index < length-1; index++ {
		swapped := false

		for j := 0; j < length-1-index; j++ {
			if cmp.Less(array[j+1], array[j]) {
				array.swap(j+1, j)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return array
}
