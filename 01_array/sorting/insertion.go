package sorting

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"
)

func InsertionSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Insertion Sort")
	if skip {
		return array
	}

	// Start from the second element, the first is already "sorted"
	for index := 1; index < length; index++ {
		insert := index
		current := array[index]

		// Shift larger elements to the right to make space for the current value
		for pass := index - 1; pass >= 0 && cmp.Less(current, array[pass]); pass-- {
			array[pass+1] = array[pass] // Shift larger elements to the right
			insert = pass               // Update the insert index to the current position
		}

		// Insert the current value at its correct position
		array[insert] = current
		helpers.PrintIteration(index-1, array)
	}

	helpers.PrintFinal(array)
	return array
}
