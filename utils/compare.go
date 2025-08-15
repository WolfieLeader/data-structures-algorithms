package utils

import "cmp"

type Ordered = cmp.Ordered
type Operator uint8

const (
	LessThan Operator = iota
	GreaterThan
	LessOrEqualTo
	GreaterOrEqualTo
	EqualTo
)

func Is[T cmp.Ordered](x T, op Operator, y T) bool {
	switch op {
	case LessThan:
		return cmp.Compare(x, y) < 0
	case GreaterThan:
		return cmp.Compare(x, y) > 0
	case LessOrEqualTo:
		return cmp.Compare(x, y) <= 0
	case GreaterOrEqualTo:
		return cmp.Compare(x, y) >= 0
	case EqualTo:
		return cmp.Compare(x, y) == 0
	default:
		panic("Invalid operator")
	}
}
