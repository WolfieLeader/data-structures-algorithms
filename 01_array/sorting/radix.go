package sorting

func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	array := copyArray(arr)

	printName("Radix Sort", array)

	// Find the maximum number to know the number of digits
	maxVal := array[0]
	for _, val := range array {
		if val > maxVal {
			maxVal = val
		}
	}

	// Perform counting sort for every digit (exp = 1, 10, 100, ...)
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countingSortByDigit(array, exp)
	}

	printFinal(array)
	return array
}

func countingSortByDigit(array []int, exp int) {
	// Create a count array for digits 0-9
	count := make([]int, 10)
	output := make([]int, len(array))

	// Count occurrences of each digit
	for _, val := range array {
		digit := (val / exp) % 10
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