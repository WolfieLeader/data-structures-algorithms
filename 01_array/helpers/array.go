package helpers

import (
	"cmp"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/array_types/dynamic"
)

type Array = dynamic.DynamicArrayType[int]

func CopyArr(arr []int, name string) (array dynamic.DynamicArrayType[int], length int, skip bool) {
	length = len(arr)

	if length <= 1 {
		return arr, length, true
	}

	array = make([]int, len(arr))
	copy(array, arr)

	PrintName(name, array)
	return array, length, false
}

func FindMaxValue(array Array) int {
	max := array[0]
	for _, value := range array {
		if cmp.Less(max, value) {
			max = value
		}
	}
	return max
}
