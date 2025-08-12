package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func RadixSort(arr []int) helpers.Array {
	array, _, skip := helpers.CopyArr(arr, "Radix Sort")
	if skip {
		return array
	}

	// Find the maximum number to know the number of digits
	maxVal := array[0]
	for _, value := range array {
		if value > maxVal {
			maxVal = value
		}
	}

	// Perform counting sort for every digit (exp = 1, 10, 100, ...)
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countingSortByDigit(array, exp)
	}

	helpers.PrintFinal(array)
	return array
}

func countingSortByDigit(array helpers.Array, exp int) {
	// Create a count array for digits 0-9
	count := make([]int, 10)
	output := make([]int, len(array))

	// Count occurrences of each digit
	for _, value := range array {
		digit := (value / exp) % 10
		count[digit]++
	}

	// Update count[i] to contain the actual position of this digit in output[]
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := len(array) - 1; i >= 0; i-- {
		digit := (array[i] / exp) % 10
		output[count[digit]-1] = array[i]
		count[digit]--
	}

	// Copy the output array to the original array
	copy(array, output)
}
