package sort

import "cmp"

func Selection[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for index := 0; index < length-1; index++ {
		minIndex := index

		for j := index + 1; j < length; j++ {
			if cmp.Less(array[j], array[minIndex]) {
				minIndex = j
			}
		}
		array.swap(index, minIndex)
	}

	return array
}
