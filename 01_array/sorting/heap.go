package sorting

func HeapSort(arr []int) Array {
	array := copyArray(arr)

	printName("Heap Sort", array)

	// Build the max heap
	for i := len(array)/2 - 1; i >= 0; i-- {
		heapify(array, len(array), i)
	}

	// Extract elements from the heap one by one
	for i := len(array) - 1; i > 0; i-- {
		array.swap(0, i)     // Move current root to end
		heapify(array, i, 0) // Call heapify on the reduced heap
	}

	printFinal(array)
	return array
}

func heapify(array Array, size int, root int) {
	largest := root // Initialize largest as root
	left := 2*root + 1
	right := 2*root + 2

	// If left child is larger than root
	if left < size && array[left] > array[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < size && array[right] > array[largest] {
		largest = right
	}

	// If largest is not root
	if largest != root {
		array.swap(root, largest)     // Swap root with largest
		heapify(array, size, largest) // Recursively heapify the affected sub-tree
	}
}
