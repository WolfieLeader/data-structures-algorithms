package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func CombSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Comb Sort")
	if skip {
		return array
	}

	gap := length
	shrunk := true

	for shrunk || gap > 1 {
		if shrunk {
			gap = max(int(float64(gap)/1.3), 1)
			shrunk = false
		}

		for index := 0; index+gap < length; index++ {
			if array[index] > array[index+gap] {
				array.Swap(index, index+gap)
				shrunk = true
			}
			helpers.PrintIteration(index, array)
		}
	}

	helpers.PrintFinal(array)
	return array
}
