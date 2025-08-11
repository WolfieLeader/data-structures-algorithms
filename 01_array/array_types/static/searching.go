package static

func (array staticArray) LinearSearch(value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func (array staticArray) BinarySearch(value int) int {
	if !array.IsSorted() {
		return -1 // Binary search requires sorted array
	}

	left, right := 0, len(array)-1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == value {
			return mid
		} else if array[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
