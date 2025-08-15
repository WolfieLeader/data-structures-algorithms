package sorting

func Heap[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for i := length/2 - 1; i >= 0; i-- {
		heapify(array, i, length)
	}

	for end := length - 1; end > 0; end-- {
		array.swap(0, end)
		heapify(array, 0, end)
	}

	return array
}

func heapify[T Ordered](array arrayType[T], root, end int) {
	for {
		left := 2*root + 1
		if left >= end {
			return
		}

		child := left
		right := left + 1

		if right < end && is(array[left], LessThan, array[right]) {
			child = right
		}

		if is(array[root], LessThan, array[child]) {
			array.swap(root, child)
			root = child
		} else {
			return
		}
	}
}
