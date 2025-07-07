package sorting

func MergeSort(arr []int) []int {
	array := copyArray(arr)

	printName("Merge Sort", array)

	array = mergeSort(array)

	printFinal(array)

	return array
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}