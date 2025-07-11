package sorting

func ShellSort(arr []int) Array {
	array, skip := initSort("Shell Sort", arr)
	if skip {
		return array
	}

	// Start with a large gap and reduce it
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i++ {
			temp := array[i]
			j := i

			// Shift elements to the right until the correct position is found
			for j >= gap && array[j-gap] > temp {
				array[j] = array[j-gap]
				j -= gap
			}
			array[j] = temp
			printIteration(i, array)
		}
	}

	printFinal(array)
	return array
}
