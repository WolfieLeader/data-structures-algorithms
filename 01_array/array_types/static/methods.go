package static

import "fmt"

type StaticArray interface {
	Replace(values ...int)
	Get(index int) (int, error)
	Set(index, value int) error
	Length() int
	Traverse()
	Clear()
	Copy() staticArray
	IndexOf(value int) int
	Swap(i, j int) error
	Reverse() staticArray
	IsSorted() bool
	LinearSearch(value int) int
	BinarySearch(value int) int
}

var _ StaticArray = (*staticArray)(nil)

func (array *staticArray) Replace(values ...int) {
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(array[:], values)
}

func (array staticArray) Get(index int) (int, error) {
	if index < 0 || index >= len(array) {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	return array[index], nil
}

func (array *staticArray) Set(index, value int) error {
	if index < 0 || index >= len(array) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	(*array)[index] = value
	return nil
}

func (array staticArray) Length() int {
	return len(array)
}

func (array staticArray) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%d, ", i, value)
	}
	fmt.Println()
}

func (array *staticArray) Clear() {
	*array = staticArray{}
}

func (array staticArray) Copy() staticArray {
	var copyArray staticArray
	copy(copyArray[:], array[:])
	return copyArray
}

func (array staticArray) IndexOf(value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func (array *staticArray) Swap(i, j int) error {
	if i < 0 || i >= len(array) || j < 0 || j >= len(array) {
		return fmt.Errorf("index out of bounds: %d or %d", i, j)
	}
	array[i], array[j] = array[j], array[i]
	return nil
}

func (array staticArray) Reverse() staticArray {
	reversed := array.Copy()
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed.Swap(left, right)
	}
	return reversed
}

func (array staticArray) IsSorted() bool {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			return false
		}
	}
	return true
}
