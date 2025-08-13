package sort

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
	out := make([]T, 0, leftLen+rightLen)

	i, j := 0, 0
	for i < leftLen && j < rightLen {
		if less(right[j], left[i]) {
			out = append(out, right[j])
			j++
		} else {
			out = append(out, left[i])
			i++
		}
	}

	out = append(out, left[i:]...)
	out = append(out, right[j:]...)
	return out
}
