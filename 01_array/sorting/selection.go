package sorting

import "fmt"

func SelectionSort(arr []int) []int {
	fmt.Println("Selection Sort:")
	fmt.Printf("Initial array: %v\n", arr)

	// The reason it's len(arr)-1 is because the last element will already be in place after the previous iterations.
	for i := range len(arr) - 1 {
		minIndex := i

		// The reason we start from i + 1 is because all elements before i are already sorted.
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element of the unsorted part
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Printf("Iteration %d: %v\n", i+1, arr)
	}
	fmt.Printf("Sorted array: %v\n", arr)

	return arr
}
