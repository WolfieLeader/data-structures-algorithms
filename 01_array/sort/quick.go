package sort

// ------ Lomuto Partition -----

func Quick[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	lomutoQuickSort(array, 0, length-1)
	return array
}

func lomutoQuickSort[T Ordered](array array[T], low int, high int) {
	if low >= high {
		return
	}
	pivot := lomutoPartition(array, low, high)
	lomutoQuickSort(array, low, pivot-1)
	lomutoQuickSort(array, pivot+1, high)
}

func lomutoPartition[T Ordered](array array[T], low int, high int) int {
	median := medianOfThree(array, low, high)
	array.swap(median, high)
	pivotVal := array[high]

	index := low
	for j := low; j < high; j++ {
		if less(array[j], pivotVal) {
			array.swap(index, j)
			index++
		}
	}

	array.swap(index, high)
	return index
}

// ------ Hoare Partition ------

func QuickHoare[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	hoareQuickSort(array, 0, length-1)
	return array
}

func hoareQuickSort[T Ordered](array array[T], low int, high int) {
	if low >= high {
		return
	}

	pivot := hoarePartition(array, low, high)
	hoareQuickSort(array, low, pivot)
	hoareQuickSort(array, pivot+1, high)
}

func hoarePartition[T Ordered](array array[T], low int, high int) int {
	median := medianOfThree(array, low, high)
	pivotVal := array[median]

	left, right := low-1, high+1
	for {
		for {
			left++
			if less(array[left], pivotVal) == false {
				break // stop on array[left] >= pivotVal
			}
		}
		for {
			right--
			if less(pivotVal, array[right]) == false {
				break // stop on array[right] <= pivotVal
			}
		}
		if left >= right {
			return right
		}
		array.swap(left, right)
	}
}

// ------ 3 Way (Dutch Flag) -----

func Quick3Way[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	threeWayQuickSort(array, 0, length-1)
	return array
}

func threeWayQuickSort[T Ordered](array array[T], low int, high int) {
	if low >= high {
		return
	}

	lesser, greater := threeWayPartition(array, low, high)
	threeWayQuickSort(array, low, lesser-1)
	threeWayQuickSort(array, greater+1, high)
}

func threeWayPartition[T Ordered](array array[T], low int, high int) (int, int) {
	median := medianOfThree(array, low, high)
	pivotVal := array[median]

	lesser, equal, greater := low, low, high
	for equal <= greater {
		switch {
		case less(array[equal], pivotVal):
			array.swap(lesser, equal)
			lesser++
			equal++
		case less(pivotVal, array[equal]):
			array.swap(equal, greater)
			greater--
		default:
			equal++
		}
	}

	return lesser, greater
}

// ------ Shared Utilities ------

func medianOfThree[T Ordered](array array[T], low, high int) int {
	mid := low + (high-low)/2
	if less(array[mid], array[low]) {
		low, mid = mid, low
	}
	if less(array[high], array[low]) {
		low, high = high, low
	}
	if less(array[high], array[mid]) {
		mid, high = high, mid
	}
	return mid // returns the median index without swapping
}
