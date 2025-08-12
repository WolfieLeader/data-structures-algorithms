package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

func GnomeSort(arr []int) helpers.Array {
	array, length, skip := helpers.CopyArr(arr, "Gnome Sort")
	if skip {
		return array
	}

	for index := 0; index < length; {
		if index == 0 || array[index] >= array[index-1] {
			index++ // Move forward if in order or at the start
		} else {
			array.Swap(index, index-1)
			index-- // Step back to check the previous element
		}
		helpers.PrintIteration(index, array)
	}

	helpers.PrintFinal(array)
	return array
}
