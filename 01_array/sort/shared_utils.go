package sort

import "cmp"

type Ordered = cmp.Ordered
type arrayType[T Ordered] []T

type operator uint8

const (
	lessThan operator = iota
	greaterThan
	lessOrEqualTo
	greaterOrEqualTo
	equalTo
)

func is[T Ordered](x T, op operator, y T) bool {
	switch op {
	case lessThan:
		return cmp.Less(x, y)
	case greaterThan:
		return cmp.Less(y, x)
	case lessOrEqualTo:
		return !cmp.Less(y, x)
	case greaterOrEqualTo:
		return !cmp.Less(x, y)
	case equalTo:
		return !cmp.Less(x, y) && !cmp.Less(y, x)
	default:
		panic("Invalid operator")
	}
}

func copyArray[T Ordered](arr []T) (arrayType[T], int) {
	length := len(arr)
	out := make(arrayType[T], length)
	copy(out, arr)
	return out, length
}

func (array *arrayType[T]) swap(i, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
