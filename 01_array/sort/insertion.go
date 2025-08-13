package sort

import "cmp"

func Insertion[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for index := 1; index < length; index++ {
		value := array[index]
		j := index - 1

		for j >= 0 && cmp.Less(value, array[j]) {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = value
	}
	return array
}
