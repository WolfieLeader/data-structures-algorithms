package sort

func Gnome[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	i := 1
	for i < length {
		if i == 0 || is(array[i], greaterOrEqualTo, array[i-1]) {
			i++
		} else {
			array.swap(i, i-1)
			i--
		}
	}

	return array
}
