package sorting

import (
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

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
		if utils.Is(value, utils.LessThan, minVal) {
			minVal = value
		}
		if utils.Is(value, utils.GreaterThan, maxVal) {
			maxVal = value
		}
	}
	return minVal, maxVal
}

type Ordered = utils.Ordered
type operator = utils.Operator

const (
	LessThan         operator = utils.LessThan
	GreaterThan      operator = utils.GreaterThan
	LessOrEqualTo    operator = utils.LessOrEqualTo
	GreaterOrEqualTo operator = utils.GreaterOrEqualTo
	EqualTo          operator = utils.EqualTo
)

func is[T Ordered](x T, op operator, y T) bool {
	return utils.Is(x, op, y)
}
