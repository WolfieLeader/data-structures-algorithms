package sort

import "cmp"

type Ordered = cmp.Ordered
type arrayType[T Ordered] []T

type operator uint8

const (
	LessThan operator = iota
	GreaterThan
	LessOrEqualTo
	GreaterOrEqualTo
	EqualTo
)

func is[T Ordered](x T, op operator, y T) bool {
	switch op {
	case LessThan:
		return cmp.Less(x, y)
	case GreaterThan:
		return cmp.Less(y, x)
	case LessOrEqualTo:
		return !cmp.Less(y, x)
	case GreaterOrEqualTo:
		return !cmp.Less(x, y)
	case EqualTo:
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
