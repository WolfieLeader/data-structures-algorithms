package main

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/01_array/sort"
	"github.com/WolfieLeader/data-structures-algorithms/01_array/structures/dynamic"
	"github.com/WolfieLeader/data-structures-algorithms/01_array/structures/matrix"
)

func main() {
	// matrixExample()
	// searchingExample()
	sortingExample()
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

func sortingExample() {
	arr := []int{0, 64, 32, 50, 50, -5, 11, 90, 25, 15000, 200, 4, 3, 50, 2, -30, 23, 0, -1, 1, 50}

	fmt.Println(sort.Bubble(arr), "- Bubble Sort")
	fmt.Println(sort.Selection(arr), "- Selection Sort")
	fmt.Println(sort.Insertion(arr), "- Insertion Sort")
	fmt.Println(sort.Merge(arr), "- Merge Sort")
	fmt.Println(sort.Quick(arr), "- Quick Sort")
	fmt.Println(sort.QuickHoare(arr), "- Quick Hoare Sort")
	fmt.Println(sort.Quick3Way(arr), "- Quick 3-Way Sort")
	fmt.Println(sort.Bucket(arr), "- Bucket Sort")
	fmt.Println(sort.Counting(arr), "- Counting Sort")
	fmt.Println(sort.Radix(arr), "- Radix Sort")
	fmt.Println(sort.Cocktail(arr), "- Cocktail Sort")
	fmt.Println(sort.Shell(arr), "- Shell Sort")
	fmt.Println(sort.Gnome(arr), "- Gnome Sort")
	fmt.Println(sort.Comb(arr), "- Comb Sort")
}
