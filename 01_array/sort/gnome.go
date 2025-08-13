package sort

func Gnome[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	index := 1
	for index < length {
		if !less(array[index], array[index-1]) {
			index++
		} else {
			array.swap(index, index-1)
			index--
		}
	}

	return array
}
