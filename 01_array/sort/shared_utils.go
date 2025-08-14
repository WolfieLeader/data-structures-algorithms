package sort

import "cmp"

type Ordered = cmp.Ordered
type arrayType[T Ordered] []T

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integers interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Numbers interface {
	Integers | Float
}

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

func (array arrayType[T]) findMinMax() (minVal T, maxVal T) {
	minVal, maxVal = array[0], array[0]
	for _, value := range array {
		if is(value, LessThan, minVal) {
			minVal = value
		}
		if is(value, GreaterThan, maxVal) {
			maxVal = value
		}
	}
	return minVal, maxVal
}
