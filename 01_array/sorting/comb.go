package sorting

func CombSort(arr []int) Array {
	array, skip := initSort("Comb Sort", arr)
	if skip {
		return array
	}

	gap := len(array)
	shrunk := true

	for shrunk || gap > 1 {
		if shrunk {
			gap = max(int(float64(gap)/1.3), 1)
			shrunk = false
		}

		for i := 0; i+gap < len(array); i++ {
			if array[i] > array[i+gap] {
				array.swap(i, i+gap)
				shrunk = true
			}
			printIteration(i, array)
		}
	}

	printFinal(array)
	return array
}
