package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func hoareQuickSort(array helpers.Array, low, high int) {
	if low < high {
		// Partition the array
		pivot := hoarePartition(array, low, high)

		// Recursively sort the left and right sub-arrays
		hoareQuickSort(array, low, pivot)
		hoareQuickSort(array, pivot+1, high)
	}
}

func hoarePartition(array helpers.Array, low, high int) int {
	pivot := array[low]
	index1 := low - 1
	index2 := high + 1

	for {

		// Move i right until we find element >= pivot
		for {
			index1++
			if array[index1] >= pivot {
				break
			}
		}

		// Move j left until we find element <= pivot
		for {
			index2--
			if array[index2] <= pivot {
				break
			}
		}

		// If pointers cross, return j
		if index1 >= index2 {
			return index2
		}

		array.Swap(index1, index2)
	}
}
