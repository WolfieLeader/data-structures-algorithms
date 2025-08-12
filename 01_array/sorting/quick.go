package sorting

import "github.com/WolfieLeader/data-structures-algorithms/01_array/helpers"

const (
	Lomuto = iota
	Hoare
	Dutch
)

func QuickSort(arr []int, algorithm int) helpers.Array {
	array, _, skip := helpers.CopyArr(arr, "Quick Sort")
	if skip {
		return array
	}

	switch algorithm {
	case Lomuto:
		helpers.PrintName("Lomuto Quick Sort", array)
		lomutoQuickSort(array, 0, len(array)-1)
	case Hoare:
		helpers.PrintName("Hoare Quick Sort", array)
		hoareQuickSort(array, 0, len(array)-1)
	case Dutch:
		helpers.PrintName("Dutch Quick Sort", array)
		dutchQuickSort(array, 0, len(array)-1)
	default:
		panic("Unknown quick sort algorithm")
	}

	helpers.PrintFinal(array)
	return array
}
