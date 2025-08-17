package main

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/dynamic"
	"github.com/WolfieLeader/data-structures-algorithms/data_structures/array/matrix"
)

func main() {
	// matrixExample()
	// searchingExample()
}

func matrixExample() {
	twoDArray := matrix.New[int](3, 3)
	twoDArray.Traverse(func(i, j int, value int) bool {
		fmt.Printf("Value at (%d, %d): %d\n", i, j, value)
		return true
	})
}

func searchingExample() {
	unsortedArray := dynamic.New(1, 23, 5, 7, 9, 11, 13, 15, 17, 19, 21, 3, 25, 27, 29, 31)
	fmt.Printf("Searching for 23 in %v, result: %d\n", unsortedArray, unsortedArray.Search(23))
	fmt.Printf("Searching for 71 in %v, result: %d\n", unsortedArray, unsortedArray.Search(71))

	sortedArray := dynamic.New(3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31)
	fmt.Printf("Searching for 15 in %v, result: %d\n", sortedArray, sortedArray.Search(15))
	fmt.Printf("Searching for 2 in %v, result: %d\n", sortedArray, sortedArray.Search(2))
}
