package sort

func Selection[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for index := 0; index < length-1; index++ {
		minIdx := index

		for j := index + 1; j < length; j++ {
			if less(array[j], array[minIdx]) {
				minIdx = j
			}
		}
		array.swap(index, minIdx)
	}

	return array
}
