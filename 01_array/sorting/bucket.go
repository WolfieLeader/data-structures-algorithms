package sorting

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"
)

func BucketSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Bucket Sort")
	if skip {
		return array
	}

	max := helpers.FindMaxValue(array)
	if max == 0 {
		helpers.PrintFinal(array)
		return array
	}

	buckets := make([][]int, length)

	// Distribute values into buckets
	for _, value := range array {
		index := value * (length - 1) / max //TODO
		buckets[index] = append(buckets[index], value)
	}

	// Sort each bucket with insertion sort
	sorted := make([]int, 0, length)

	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sorted = append(sorted, insertionSort(bucket)...)
		}
	}

	helpers.PrintFinal(sorted)
	return sorted
}

func insertionSort(array helpers.Array) helpers.Array {
	for index := 1; index < len(array); index++ {
		insert := index
		current := array[index]

		for pass := index - 1; pass >= 0 && cmp.Less(current, array[pass]); pass-- {
			array[pass+1] = array[pass]
			insert = pass
		}

		array[insert] = current
	}

	return array
}
