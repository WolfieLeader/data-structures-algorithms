package sorting

// InsertionSort builds a sorted array one item at a time by inserting
// each element into its correct position within the sorted portion.
func InsertionSort(array []int) []int {
	printName("Insertion Sort", array)

	// Start from the second element, the first is already "sorted"
	for i := 1; i < len(array); i++ {
		insertIndex := i
		current := array[i]

		// Shift larger elements to the right to make space for the current value
		for j := i - 1; j >= 0 && (array[j] > current); j-- {
			array[j+1] = array[j] // Shift larger elements to the right
			insertIndex = j       // Update the insert index to the current position
		}

		// Insert the current value at its correct position
		array[insertIndex] = current
		printIteration(i-1, array)
	}

	printFinal(array)
	return array
}
