package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func MergeSort(arr []int) helpers.Array {
	array, _, skip := helpers.CopyArr(arr, "Merge Sort")
	if skip {
		return array
	}

	array = mergeSort(array)

	helpers.PrintFinal(array)
	return array
}

func mergeSort(array helpers.Array) helpers.Array {
	length := len(array)
	if length <= 1 {
		return array
	}

	mid := length / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(left, right helpers.Array) helpers.Array {
	result := make([]int, 0, len(left)+len(right))
	index1, index2 := 0, 0

	for index1 < len(left) && index2 < len(right) {
		if left[index1] <= right[index2] {
			result = append(result, left[index1])
			index1++
		} else {
			result = append(result, right[index2])
			index2++
		}
	}

	// Append any remaining elements
	result = append(result, left[index1:]...)
	result = append(result, right[index2:]...)

	return result
}
