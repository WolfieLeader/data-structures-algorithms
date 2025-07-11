package sorting

func hoareQuickSort(array Array, low, high int) {
	if low < high {
		// Partition the array
		pivot := hoarePartition(array, low, high)

		// Recursively sort the left and right sub-arrays
		hoareQuickSort(array, low, pivot)
		hoareQuickSort(array, pivot+1, high)
	}
}

func hoarePartition(array Array, low, high int) int {
	pivot := array[low]
	i := low - 1
	j := high + 1

	for {

		// Move i right until we find element >= pivot
		for {
			i++
			if array[i] >= pivot {
				break
			}
		}

		// Move j left until we find element <= pivot
		for {
			j--
			if array[j] <= pivot {
				break
			}
		}

		// If pointers cross, return j
		if i >= j {
			return j
		}

		array.swap(i, j)
	}
}
