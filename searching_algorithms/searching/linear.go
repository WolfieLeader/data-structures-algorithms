package searching

import (
	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func LinearSearch[T utils.Ordered](array []T, value T) int {
	for i, v := range array {
		if utils.Is(v, utils.EqualTo, value) {
			return i
		}
	}
	return -1
}
