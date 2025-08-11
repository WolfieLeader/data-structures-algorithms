package dynamic

type dynamicArray []int

func NewDynamicArray(values ...int) dynamicArray {
	return append(dynamicArray{}, values...)
}
