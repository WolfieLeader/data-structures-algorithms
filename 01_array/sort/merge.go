package sort

import "cmp"

func Merge[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	mid := length / 2
	left := Merge(array[:mid])
	right := Merge(array[mid:])
	return merge(left, right)
}

func merge[T cmp.Ordered](left []T, right []T) []T {
	leftLength, rightLength := len(left), len(right)
	out := make([]T, 0, leftLength+rightLength)
	i, j := 0, 0

	for i < leftLength && j < rightLength {
		if cmp.Less(right[j], left[i]) {
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
