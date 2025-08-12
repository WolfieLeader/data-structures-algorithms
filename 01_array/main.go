package main

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/array_types/dynamic"
	"github.com/WolfieLeader/data-structures-algorithms/01_array/array_types/matrix"
	"github.com/WolfieLeader/data-structures-algorithms/01_array/sorting"
)

func main() {
	matrixExample()
	searchingExample()
	sortingExample()
}

func matrixExample() {
	twoDArray, err := matrix.New(3, 3, []int{1, 2, 3}, []int{4})
	if err != nil {
		fmt.Println("Error creating two-dimensional array:", err)
		return
	}

	twoDArray.Traverse()
}

func searchingExample() {
	unsortedArray := dynamic.New(1, 23, 5, 7, 9, 11, 13, 15, 17, 19, 21, 3, 25, 27, 29, 31)
	fmt.Printf("Searching for 23 in %v, result: %d\n", unsortedArray, unsortedArray.LinearSearch(23))
	fmt.Printf("Searching for 71 in %v, result: %d\n", unsortedArray, unsortedArray.LinearSearch(71))

	sortedArray := dynamic.New(3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31)
	fmt.Printf("Searching for 15 in %v, result: %d\n", sortedArray, sortedArray.LinearSearch(15))
	fmt.Printf("Searching for 2 in %v, result: %d\n", sortedArray, sortedArray.LinearSearch(2))
}

func sortingExample() {
	arr := dynamic.New(0, 100, 64, 34, 25, 25, 12, 22, 11, 90, 25, 301, 299, 4, 3, 2, 23)

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
