package sort

func Bubble[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for pass := 0; pass < length-1; pass++ {
		swapped := false

		for i := 0; i < length-1-pass; i++ {
			if is(array[i], GreaterThan, array[i+1]) {
				array.swap(i+1, i)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return array
}
