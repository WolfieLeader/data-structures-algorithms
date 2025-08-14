package sort

func Selection[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for pass := 0; pass < length-1; pass++ {
		minIdx := pass

		for i := pass + 1; i < length; i++ {
			if is(array[i], lessThan, array[minIdx]) {
				minIdx = i
			}
		}
		array.swap(pass, minIdx)
	}

	return array
}
