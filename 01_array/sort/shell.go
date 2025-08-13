package sort

func Shell[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := array[i]
			j := i

			for j >= gap && less(temp, array[j-gap]) {
				array[j] = array[j-gap]
				j -= gap
			}
			array[j] = temp
		}
		gap /= 2
	}

	return array
}
