package searching

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // Return the index if found
		}
	}
	return -1 // Return -1 if not found
}
