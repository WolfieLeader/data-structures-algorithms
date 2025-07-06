package sorting

// BubbleSort compares and swaps adjacent elements until the array is sorted
func BubbleSort(arr []int) []int {
	array := copyArray(arr)
	
	printName("Bubble Sort", array)

	// n-1 passes are enough, each pass places one value in its final spot
	for i := range len(array) - 1 {
		swapped := false

		// "bubble up" the largest value to the end of the unsorted part
		for j := range len(array) - 1 - i {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
		printIteration(i, array)
	}

	printFinal(array)
	return array
}
