package sorting

func CountingSort(arr []int) Array {
	array, skip := initSort("Counting Sort", arr)
	if skip {
		return array
	}

	// Find the max value in the array
	max := array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
	}

	// Create a count array of size maxVal + 1
	count := make([]int, max+1)

	// Count occurrences of each number
	for _, value := range array {
		count[value]++
	}

	// Reconstruct the sorted array
	sorted := make([]int, 0, len(array))
	for num, freq := range count {
		for range freq {
			sorted = append(sorted, num)
		}
	}

	printFinal(sorted)
	return sorted
}
