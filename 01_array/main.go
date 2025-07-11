package main

import (
	"github.com/WolfieLeader/data-structures-algorithms/01_array/sorting"
)

// Array is a continuous block of memory that stores elements of the same type.

func main() {
	// arrays.StaticArrayExample()
	// arrays.DynamicArrayExample()
	// arrays.MultiDimensionalArrayExample()

	arr := []int{0, 100, 64, 34, 25, 25, 12, 22, 11, 90, 25, 301}

	// sorting.BubbleSort(arr)
	// sorting.SelectionSort(arr)
	// sorting.InsertionSort(arr)
	// sorting.GnomeSort(arr)
	sorting.CocktailSort(arr)
	// sorting.QuickSort(arr,sorting.Lomuto)
	// sorting.QuickSort(arr,sorting.Hoare)
	// sorting.QuickSort(arr,sorting.Dutch)
	// sorting.CountingSort(arr)
	// sorting.RadixSort(arr)
	// sorting.MergeSort(arr)
}
