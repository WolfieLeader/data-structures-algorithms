package main

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/searching"
	"github.com/WolfieLeader/data-structures-algorithms/01_array/sorting"
)

func main() {
	// arrays.StaticArrayExample()
	// arrays.DynamicArrayExample()
	// arrays.MultiDimensionalArrayExample()
	// searchingExample()
	// sortingExample()
}

func searchingExample() {
	unsortedArray := []int{1, 23, 5, 7, 9, 11, 13, 15, 17, 19, 21, 3, 25, 27, 29, 31}
	fmt.Printf("Searching for 23 in %v, result: %d\n", unsortedArray, searching.LinearSearch(unsortedArray, 23))
	fmt.Printf("Searching for 71 in %v, result: %d\n", unsortedArray, searching.LinearSearch(unsortedArray, 71))

	sortedArray := []int{3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	fmt.Printf("Searching for 15 in %v, result: %d\n", sortedArray, searching.LinearSearch(sortedArray, 15))
	fmt.Printf("Searching for 2 in %v, result: %d\n", sortedArray, searching.LinearSearch(sortedArray, 2))
}

func sortingExample() {
	arr := []int{0, 100, 64, 34, 25, 25, 12, 22, 11, 90, 25, 301, 299, 4, 3, 2, 23}

	sorting.BubbleSort(arr)
	// sorting.SelectionSort(arr)
	// sorting.InsertionSort(arr)
	// sorting.GnomeSort(arr)
	// sorting.CocktailSort(arr)
	// sorting.CombSort(arr)
	// sorting.ShellSort(arr)
	// sorting.MergeSort(arr)
	// sorting.QuickSort(arr, sorting.Lomuto)
	// sorting.QuickSort(arr,sorting.Hoare)
	// sorting.QuickSort(arr,sorting.Dutch)
	// sorting.HeapSort(arr)
	// sorting.CountingSort(arr)
	// sorting.RadixSort(arr)
	// sorting.BucketSort(arr)
}
