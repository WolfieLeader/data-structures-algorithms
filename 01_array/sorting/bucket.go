package sorting

func BucketSort(arr []int) Array {
	array, skip := initSort("Bucket Sort", arr)
	if skip {
		return array
	}

	max := findMax(array)
	if max == 0 {
		printFinal(array)
		return array
	}

	numBuckets := len(array)
	buckets := make([][]int, numBuckets)

	// Distribute values into buckets
	for _, value := range array {
		index := value * (numBuckets - 1) / max
		buckets[index] = append(buckets[index], value)
	}

	// Sort each bucket with insertion sort
	sorted := make([]int, 0, len(array))
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sorted = append(sorted, insertionSort(bucket)...)
		}
	}

	printFinal(sorted)
	return sorted
}

func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func insertionSort(array Array) Array {
	for i := 1; i < len(array); i++ {
		insert := i
		current := array[i]

		for j := i - 1; j >= 0 && (array[j] > current); j-- {
			array[j+1] = array[j]
			insert = j
		}

		array[insert] = current
	}

	return array
}
