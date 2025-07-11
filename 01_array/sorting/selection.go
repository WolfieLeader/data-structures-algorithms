package sorting

// SelectionSort repeatedly finds the smallest value and moves it to the front
func SelectionSort(arr []int) Array {
	array, skip := initSort("Selection Sort", arr)
	if skip {
		return array
	}

	// Only need n-1 passes; the last element will already be sorted
	for i := range len(array) - 1 {
		min := i

		// "Select" the smallest value in the unsorted part
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		array.swap(i, min)
		printIteration(i, array)
	}

	printFinal(array)
	return array
}
