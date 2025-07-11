package sorting

func CocktailSort(arr []int) Array {
	array := copyArray(arr)

	printName("Cocktail Sort", array)

	swapped := true
	start := 0
	end := len(array) - 1

	for swapped {
		swapped = false

		// Forward pass
		for i := start; i < end; i++ {
			// Normal bubble sort logic
			if array[i] > array[i+1] {
				array.swap(i, i+1)
				swapped = true
			}
			printIteration(i, array)
		}

		// If no elements were swapped, the array is sorted
		if !swapped {
			break
		}

		swapped = false
		end--

		// Backward pass
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array.swap(i, i+1)
				swapped = true
			}
			printIteration(i, array)
		}

		start++
	}

	printFinal(array)
	return array
}
