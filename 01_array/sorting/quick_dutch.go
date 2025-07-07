package sorting

func dutchQuickSort(array []int, low, high int) {
	if low >= 0 && low < high {
		lt, gt := dutchPartition(array, low, high)

		// Recursively sort the left(smaller) and right(bigger) sub-arrays and not the equal sub-array
		dutchQuickSort(array, low, lt-1)
		dutchQuickSort(array, gt+1, high)
	}
}

func dutchPartition(array []int, low, high int) (int, int) {
	pivot := array[(low+high)/2] // Choose the middle element as pivot
	lt, eq, gt := low, low, high // Lesser, equal and greater index

	for eq <= gt {
		if array[eq] < pivot {
			// Swap lesser element with current element
			array[lt], array[eq] = array[eq], array[lt]
			lt++
			eq++
		} else if array[eq] > pivot {
			// Swap greater element with current element
			array[eq], array[gt] = array[gt], array[eq]
			gt--
		} else {
			// If equal, just move to the next element
			eq++
		}
	}

	return lt, gt
}