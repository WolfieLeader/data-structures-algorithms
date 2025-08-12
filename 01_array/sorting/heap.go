package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func HeapSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Heap Sort")
	if skip {
		return array
	}

	// Build the max heap
	for index := (length / 2) - 1; index >= 0; index-- {
		heapify(array, length, index)
	}

	// Extract elements from the heap one by one
	for i := length - 1; i > 0; i-- {
		array.Swap(0, i)     // Move current root to end
		heapify(array, i, 0) // Call heapify on the reduced heap
	}

	helpers.PrintFinal(array)
	return array
}

//TODO
func heapify(array helpers.Array, size int, root int) {
	max := root // Initialize max as root
	l := 2*root + 1
	r := 2*root + 2

	// If left child is larger than root
	if l < size && array[l] > array[max] {
		max = l
	}

	// If right child is larger than max so far
	if r < size && array[r] > array[max] {
		max = r
	}

	// If max is not root
	if max != root {
		array.Swap(root, max)     // Swap root with max
		heapify(array, size, max) // Recursively heapify the affected sub-tree
	}
}
