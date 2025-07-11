package sorting

func CocktailSort(arr []int) []int {
	array := copyArray(arr)

	printName("Cocktail Sort", array)

	n := len(array)
	swapped := true
	start := 0
	end := n - 1

	for swapped {
		swapped = false

		// Forward pass
		for i := start; i < end; i++ {
			// Normal bubble sort logic
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
			printIteration(i, array)
		}

		if !swapped {
			break
		}

		swapped = false
		end--

		// Backward pass
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
			printIteration(i, array)
		}

		start++
	}

	printFinal(array)
	return array
}
