package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func lomutoQuickSort(array helpers.Array, low, high int) {
	if low < high {

		// Partition the array
		pivot := lomutoPartition(array, low, high)

		// Recursively sort the left and right sub-arrays
		lomutoQuickSort(array, low, pivot-1)
		lomutoQuickSort(array, pivot+1, high)
	}
}

func lomutoPartition(array helpers.Array, low, high int) int {
	medianOfThree(array, low, high)

	pivot := array[high] // Choose the last element as pivot
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++

			array.Swap(i, j)
		}
	}

	// Place the pivot at the correct position, since it stayed at the end
	array.Swap(i+1, high)
	return i + 1
}

func medianOfThree(array helpers.Array, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	// Swap if low is greater than mid
	if array[low] > array[mid] {
		array.Swap(low, mid)
	}

	// Swap if low is greater than high
	if array[low] > array[high] {
		array.Swap(low, high)
	}

	// Swap if high is greater than mid, ensuring the pivot(high) is the median
	if array[mid] < array[high] {
		array.Swap(mid, high)
	}
}
