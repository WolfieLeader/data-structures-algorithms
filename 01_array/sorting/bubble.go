package sorting

import "fmt"

func BubbleSort(arr []int) []int {
	fmt.Println("Bubble Sort:")
	fmt.Printf("Initial array: %v\n", arr)

	for i := range len(arr) - 1 {
		// Track if a swap was made in this iteration
		swapped := false

		// Last i elements are already sorted, so we can skip them
		for j := range (len(arr) - 1) - i {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no swaps were made, the array is sorted
		if !swapped {
			break
		}

		fmt.Printf("Iteration %d: %v\n", i+1, arr)
	}

	fmt.Printf("Sorted array: %v\n", arr)
	return arr
}
