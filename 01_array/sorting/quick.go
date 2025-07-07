package sorting

const (
	Lomuto = iota
	Hoare
)

func QuickSort(arr []int, algorithm int) []int {
	array := copyArray(arr)

	switch algorithm {
	case Lomuto:
		printName("Lomuto Quick Sort", array)
		lomutoQuickSort(array, 0, len(array)-1)
	case Hoare:
		printName("Hoare Quick Sort", array)
		hoareQuickSort(array, 0, len(array)-1)
	default:
		panic("Unknown quick sort algorithm")
	}

	printFinal(array)
	return array
}