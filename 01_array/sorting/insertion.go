package sorting

func InsertionSort(arr []int) Array {
	array, skip := initSort("Insertion Sort", arr)
	if skip {
		return array
	}

	// Start from the second element, the first is already "sorted"
	for i := 1; i < len(array); i++ {
		insert := i
		current := array[i]

		// Shift larger elements to the right to make space for the current value
		for j := i - 1; j >= 0 && (array[j] > current); j-- {
			array[j+1] = array[j] // Shift larger elements to the right
			insert = j            // Update the insert index to the current position
		}

		// Insert the current value at its correct position
		array[insert] = current
		printIteration(i-1, array)
	}

	printFinal(array)
	return array
}
