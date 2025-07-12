package array_types

import (
	"errors"
	"fmt"
)

type DynamicArray []int

func NewDynamicArray(values ...int) DynamicArray {
	return append(DynamicArray{}, values...)
}

func (array *DynamicArray) Replace(values ...int) {
	*array = append((*array)[:0], values...)
}

func (array DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= len(array) {
		return 0, errors.New("index out of bounds")
	}
	return array[index], nil
}

func (array *DynamicArray) Set(index, value int) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*array)[index] = value
	return nil
}

func (array *DynamicArray) Insert(index, value int) error {
	if index < 0 || index > len(*array) {
		return errors.New("index out of bounds")
	}
	*array = append(*array, 0)                 // grow by 1
	copy((*array)[index+1:], (*array)[index:]) // shift right
	(*array)[index] = value

	return nil
}

func (array *DynamicArray) Append(value ...int) {
	*array = append(*array, value...)
}

func (array *DynamicArray) Delete(index int) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	*array = append((*array)[:index], (*array)[index+1:]...)
	return nil
}

func (array DynamicArray) Length() int {
	return len(array)
}

func (array DynamicArray) Capacity() int {
	return cap(array)
}

func (array DynamicArray) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%d, ", i, value)
	}
	fmt.Println()
}

func (array *DynamicArray) Clear() {
	*array = DynamicArray{}
}

func (array DynamicArray) Copy() DynamicArray {
	newArray := make(DynamicArray, len(array))
	copy(newArray, array)
	return newArray
}

func (array DynamicArray) Slice(start, end int) (DynamicArray, error) {
	if start < 0 || end > len(array) || start >= end {
		return nil, errors.New("invalid slice range")
	}
	slicedData := make(DynamicArray, end-start)
	copy(slicedData, array[start:end])
	return slicedData, nil
}

func (array DynamicArray) IndexOf(value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func (array *DynamicArray) Swap(i, j int) error {
	if i < 0 || i >= len(*array) || j < 0 || j >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	return nil
}

func (array DynamicArray) Reverse() DynamicArray {
	reversed := array.Copy()
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed.Swap(left, right)
	}
	return reversed
}
