package sorting

func Merge[T Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	mid := length / 2
	left := Merge(array[:mid])
	right := Merge(array[mid:])
	return merge(left, right)
}

func merge[T Ordered](left []T, right []T) []T {
	leftLen, rightLen := len(left), len(right)
	array := make([]T, 0, leftLen+rightLen)

	i, j := 0, 0
	for i < leftLen && j < rightLen {
		if is(left[i], GreaterThan, right[j]) {
			array = append(array, right[j])
			j++
		} else {
			array = append(array, left[i])
			i++
		}
	}

	array = append(array, left[i:]...)
	array = append(array, right[j:]...)
	return array
}
