package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func CountingSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Counting Sort")
	if skip {
		return array
	}

	max := helpers.FindMaxValue(array)
	count := make([]int, max+1)

	// Count occurrences of each number
	for _, value := range array {
		count[value]++
	}

	// Reconstruct the sorted array
	sorted := make([]int, 0, length)

	for num, freq := range count {
		for range freq {
			sorted = append(sorted, num)
		}
	}

	helpers.PrintFinal(sorted)
	return sorted
}
