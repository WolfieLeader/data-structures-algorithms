package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func CocktailSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Cocktail Sort")
	if skip {
		return array
	}

	swapped := true
	start := 0
	end := length - 1

	for swapped {
		swapped = false

		// Forward pass
		for index1 := start; index1 < end; index1++ {
			// Normal bubble sort logic
			if array[index1] > array[index1+1] {
				array.Swap(index1, index1+1)
				swapped = true
			}
			helpers.PrintIteration(index1, array)
		}

		// If no elements were swapped, the array is sorted
		if !swapped {
			break
		}

		swapped = false
		end--

		// Backward pass
		for index2 := end - 1; index2 >= start; index2-- {
			if array[index2] > array[index2+1] {
				array.Swap(index2, index2+1)
				swapped = true
			}
			helpers.PrintIteration(index2, array)
		}

		start++
	}

	helpers.PrintFinal(array)
	return array
}
