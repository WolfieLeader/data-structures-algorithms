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
	twoDArray, err := matrix.New(3, 3, []int{1, 2, 3}, []int{4})
	if err != nil {
		fmt.Println("Error creating two-dimensional array:", err)
		return
	}

	twoDArray.Traverse()
}

func searchingExample() {
	unsortedArray := dynamic.New(1, 23, 5, 7, 9, 11, 13, 15, 17, 19, 21, 3, 25, 27, 29, 31)
	fmt.Printf("Searching for 23 in %v, result: %d\n", unsortedArray, unsortedArray.LinearSearch(23))
	fmt.Printf("Searching for 71 in %v, result: %d\n", unsortedArray, unsortedArray.LinearSearch(71))

	sortedArray := dynamic.New(3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31)
	fmt.Printf("Searching for 15 in %v, result: %d\n", sortedArray, sortedArray.LinearSearch(15))
	fmt.Printf("Searching for 2 in %v, result: %d\n", sortedArray, sortedArray.LinearSearch(2))
}
