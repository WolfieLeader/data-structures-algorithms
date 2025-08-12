package sorting

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"
)

func ShellSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Shell Sort")
	if skip {
		return array
	}

	// Start with a large gap and reduce it
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			temp := array[i]
			j := i

			// Shift elements to the right until the correct position is found
			for j >= gap && cmp.Less(temp, array[j-gap]) {
				array[j] = array[j-gap]
				j -= gap
			}
			array[j] = temp
			helpers.PrintIteration(i, array)
		}
	}

	helpers.PrintFinal(array)
	return array
}
