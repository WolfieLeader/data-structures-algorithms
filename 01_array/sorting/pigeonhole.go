package sorting

func PigeonholeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	array := copyArray(arr)
	printName("Pigeonhole Sort", array)

	// Find the min and max
	min, max := array[0], array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	// Create pigeonholes
	size := max - min + 1
	holes := make([][]int, size) // slice of slices

	// Place elements in pigeonholes
	for _, value := range array {
		i := value - min
		holes[i] = append(holes[i], value)
	}

	// Collect elements from pigeonholes
	i := 0
	for _, bucket := range holes {
		for _, value := range bucket {
			array[i] = value
			i++
		}
	}

	printFinal(array)
	return array
}
