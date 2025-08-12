package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func dutchQuickSort(array helpers.Array, low, high int) {
	if low >= 0 && low < high {
		lesser, greater := dutchPartition(array, low, high)

		// Recursively sort the left(smaller) and right(bigger) sub-arrays and not the equal sub-array
		dutchQuickSort(array, low, lesser-1)
		dutchQuickSort(array, greater+1, high)
	}
}

func dutchPartition(array helpers.Array, low, high int) (int, int) {
	pivot := array[(low+high)/2]             // Choose the middle element as pivot
	lesser, equal, greater := low, low, high // Lesser, equal and greater index

	for equal <= greater {
		if array[equal] < pivot {
			// Swap lesser element with current element
			array.Swap(equal, lesser)
			lesser++
			equal++
		} else if array[equal] > pivot {
			// Swap greater element with current element
			array.Swap(equal, greater)
			greater--
		} else {
			// If equal, just move to the next element
			equal++
		}
	}

	return lesser, greater
}
