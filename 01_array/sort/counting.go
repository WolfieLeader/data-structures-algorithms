package sort

func Counting[T Integers](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	minVal, maxVal := array[0], array[0]
	for _, value := range array {
		if is(value, LessThan, minVal) {
			minVal = value
		}
		if is(value, GreaterThan, maxVal) {
			maxVal = value
		}
	}

	rangeSize := int(maxVal-minVal) + 1
	count := make([]int, rangeSize)

	for _, value := range array {
		count[int(value-minVal)]++
	}

	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}

	out := make([]T, length)
	for i := length - 1; i >= 0; i-- {
		value := array[i]
		index := int(value - minVal)
		count[index]--
		out[count[index]] = value
	}

	return out
}
