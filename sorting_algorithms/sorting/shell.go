package sorting

func Shell[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	gap := length / 2
	for gap > 0 {
		for pass := gap; pass < length; pass++ {
			temp := array[pass]
			i := pass

			for i >= gap && is(temp, LessThan, array[i-gap]) {
				array[i] = array[i-gap]
				i -= gap
			}
			array[i] = temp
		}
		gap /= 2
	}

	return array
}
