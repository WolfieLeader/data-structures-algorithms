package sorting

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"
)

// SelectionSort repeatedly finds the smallest value and moves it to the front
func SelectionSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Selection Sort")
	if skip {
		return array
	}

	// Only need n-1 passes; the last element will already be sorted
	for index := range length - 1 {
		min := index

		// "Select" the smallest value in the unsorted part
		for pass := index + 1; pass < length; pass++ {
			if cmp.Less(array[pass], array[min]) {
				min = pass
			}
		}

		array.Swap(index, min)
		helpers.PrintIteration(index, array)
	}

	helpers.PrintFinal(array)
	return array
}
