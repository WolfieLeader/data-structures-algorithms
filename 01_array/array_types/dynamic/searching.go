package dynamic

import "cmp"

func (array DynamicArrayType[T]) LinearSearch(value T) int {
	for i, v := range array {
		if cmp.Compare(v, value) == 0 {
			return i
		}
	}
	return -1
}

func (array DynamicArrayType[T]) BinarySearch(value T) int {
	if !array.IsSorted() {
		return -1 // Binary search requires sorted array
	}

	for l, r := 0, len(array)-1; l <= r; {
		mid := (l + r) / 2

		switch cmp.Compare(array[mid], value) {
		case 0:
			return mid
		case -1:
			l = mid + 1
		case 1:
			r = mid - 1
		}
	}

	return -1
}
