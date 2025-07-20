package array_types

import (
	"errors"
	"fmt"
)

type dynamicArray []int

type DynamicArray interface {
	Replace(values ...int)
	Get(index int) (int, error)
	Set(index, value int) error
	Insert(index, value int) error
	Append(value ...int)
	Prepend(value ...int)
	Delete(index int) error
	Length() int
	Capacity() int
	Traverse()
	Clear()
	Copy() dynamicArray
	Slice(start, end int) (dynamicArray, error)
	IndexOf(value int) int
	Swap(i, j int) error
	Reverse() dynamicArray
}

var _ DynamicArray = (*dynamicArray)(nil)

func NewDynamicArray(values ...int) dynamicArray {
	return append(dynamicArray{}, values...)
}

func (array *dynamicArray) Replace(values ...int) {
	*array = append((*array)[:0], values...)
}

func (array dynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= len(array) {
		return 0, errors.New("index out of bounds")
	}
	return array[index], nil
}

func (array *dynamicArray) Set(index, value int) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer but good practice to use pointer receiver for mutating methods
	(*array)[index] = value
	return nil
}

func (array *dynamicArray) Insert(index, value int) error {
	if index < 0 || index > len(*array) {
		return errors.New("index out of bounds")
	}
	*array = append(*array, 0)                 // grow by 1
	copy((*array)[index+1:], (*array)[index:]) // shift right
	(*array)[index] = value

	return nil
}

func (array *dynamicArray) Append(value ...int) {
	*array = append(*array, value...)
}

func (array *dynamicArray) Prepend(value ...int) {
	*array = append(value, *array...)
}

func (array *dynamicArray) Delete(index int) error {
	if index < 0 || index >= len(*array) {
		return errors.New("index out of bounds")
	}
	*array = append((*array)[:index], (*array)[index+1:]...)
	return nil
}

func (array dynamicArray) Length() int {
	return len(array)
}

func (array dynamicArray) Capacity() int {
	return cap(array)
}

func (array dynamicArray) Traverse() {
	for i, value := range array {
		fmt.Printf("[%d]:%d, ", i, value)
	}
	fmt.Println()
}

func (array *dynamicArray) Clear() {
	*array = dynamicArray{}
}

func (array dynamicArray) Copy() dynamicArray {
	newArray := make(dynamicArray, len(array))
	copy(newArray, array)
	return newArray
}

func (array dynamicArray) Slice(start, end int) (dynamicArray, error) {
	if start < 0 || end > len(array) || start >= end {
		return nil, errors.New("invalid slice range")
	}
	slicedData := make(dynamicArray, end-start)
	copy(slicedData, array[start:end])
	return slicedData, nil
}

func (array dynamicArray) IndexOf(value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func (array *dynamicArray) Swap(i, j int) error {
	if i < 0 || i >= len(*array) || j < 0 || j >= len(*array) {
		return errors.New("index out of bounds")
	}
	// No need for pointer receiver since a slice is a reference type to the underlying array but good practice
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	return nil
}

func (array dynamicArray) Reverse() dynamicArray {
	reversed := array.Copy()
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}
	return reversed
}
