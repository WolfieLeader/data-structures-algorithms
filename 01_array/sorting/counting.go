package sorting

func CountingSort(arr []int) Array {
	printName("Counting Sort", arr)

	if len(arr) == 0 {
		return arr
	}

	// Find the max value in the array
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	// Create a count array of size maxVal + 1
	count := make([]int, maxVal+1)

	// Count occurrences of each number
	for _, val := range arr {
		count[val]++
	}

	// Reconstruct the sorted array
	sorted := make([]int, 0, len(arr))
	for num, freq := range count {
		for range freq {
			sorted = append(sorted, num)
		}
	}

	printFinal(sorted)

	return sorted
}
