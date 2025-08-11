package static

const SIZE = 5

type staticArray [SIZE]int

func NewStaticArray(values ...int) staticArray {
	var staticArray staticArray
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(staticArray[:], values)
	return staticArray
}
