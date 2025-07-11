package sorting

func GnomeSort(arr []int) []int {
	array := copyArray(arr)

	printName("Gnome Sort", array)

	i := 0
	for i < len(array) {
		if i == 0 || array[i] >= array[i-1] {
			i++ // Move forward if in order or at the start
		} else {
			array[i], array[i-1] = array[i-1], array[i] // Swap if out of order
			i--                                         // Step back to check the previous element
		}
		printIteration(i, array)
	}

	printFinal(array)
	return array
}
