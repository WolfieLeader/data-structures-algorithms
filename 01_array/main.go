package main

import (
	"github.com/WolfieLeader/data-structures-algorithms/01_array/sorting"
)

// Array is a continuous block of memory that stores elements of the same type.

func main() {
	// arrays.StaticArrayExample()
	// arrays.DynamicArrayExample()
	// arrays.MultiDimensionalArrayExample()

	arr := []int{0, 100, 64, 34, 25, 12, 22, 11, 90, 301}

	// sorting.BubbleSort(arr)
	// sorting.SelectionSort(arr)
	// sorting.InsertionSort(arr)
	sorting.QuickSort(arr)
}
