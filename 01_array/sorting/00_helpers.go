package sorting

import "fmt"

func printName(name string, array []int) { fmt.Printf("\n%s:\nInitial array: %v\n", name, array) }
func printIteration(i int, array []int)  { fmt.Printf("- Iteration %d: %v\n", i+1, array) }
func printFinal(array []int)             { fmt.Printf("Sorted array: %v\n", array) }

type Array []int

func copyArray(arr []int) Array {
	copied := make([]int, len(arr))
	copy(copied, arr)
	return copied
}

func (a Array) swap(i, j int) {
	if i < 0 || j < 0 || i >= len(a) || j >= len(a) {
		return // Prevent out-of-bounds access
	}
	a[i], a[j] = a[j], a[i]
}
