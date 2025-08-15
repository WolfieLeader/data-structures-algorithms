package searching

import "github.com/WolfieLeader/data-structures-algorithms/utils"

func BinarySearch[T utils.Ordered](array []T, value T) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := low + (high-low)/2

		if utils.Is(array[mid], utils.EqualTo, value) {
			return mid
		}

		if utils.Is(array[mid], utils.LessThan, value) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
