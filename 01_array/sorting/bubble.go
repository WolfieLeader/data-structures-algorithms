package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

// BubbleSort compares and swaps adjacent elements until the array is sorted
func BubbleSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Bubble Sort")
	if skip {
		return array
	}

	// n-1 passes are enough, each pass places one value in its final spot
	for index := range length - 1 {
		swapped := false

		// "bubble up" the largest value to the end of the unsorted part
		for pos := range length - 1 - index {
			if array[pos] > array[pos+1] {
				array.Swap(pos, pos+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}

		helpers.PrintIteration(index, array)
	}

	helpers.PrintFinal(array)
	return array
}
