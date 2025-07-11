package sorting

func GnomeSort(arr []int) Array {
	array, skip := initSort("Gnome Sort", arr)
	if skip {
		return array
	}

	for i := 0; i < len(array); {
		if i == 0 || array[i] >= array[i-1] {
			i++ // Move forward if in order or at the start
		} else {
			array.swap(i, i-1)
			i-- // Step back to check the previous element
		}
		printIteration(i, array)
	}

	printFinal(array)
	return array
}
