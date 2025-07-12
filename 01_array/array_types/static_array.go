package array_types

import "fmt"

const SIZE = 5

type StaticArray [SIZE]int

func NewStaticArray(values ...int) StaticArray {
	var staticArray StaticArray
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(staticArray[:], values)
	return staticArray
}

func (array *StaticArray) Replace(values ...int) {
	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(array[:], values)
}

func (array StaticArray) Get(index int) (int, error) {
	if index < 0 || index >= len(array) {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	return array[index], nil
}

func (array *StaticArray) Set(index, value int) error {
	if index < 0 || index >= len(array) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	(*array)[index] = value
	return nil
}

func (array StaticArray) Length() int {
	return len(array)
}

func (array StaticArray) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%d, ", i, value)
	}
	fmt.Println()
}

func (array *StaticArray) Clear() {
	*array = StaticArray{}
}

func (array StaticArray) Copy() StaticArray {
	var copyArray StaticArray
	copy(copyArray[:], array[:])
	return copyArray
}

func (array StaticArray) IndexOf(value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func (array *StaticArray) Swap(i, j int) error {
	if i < 0 || i >= len(array) || j < 0 || j >= len(array) {
		return fmt.Errorf("index out of bounds: %d or %d", i, j)
	}
	array[i], array[j] = array[j], array[i]
	return nil
}

func (array StaticArray) Reverse() StaticArray {
	reversed := array.Copy()
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed.Swap(left, right)
	}
	return reversed
}
