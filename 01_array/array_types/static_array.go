package array_types

import "fmt"

func StaticArrayExample() {
	fmt.Println("\nStatic Array Example:")

	//Static staticArr of size 5
	var staticArr [5]int

	// Initialize the static array with values
	staticArr = [...]int{1, 2, 3, 4, 5}

	// Assign a value at index X (O(1) time complexity)
	staticArr[4] = 1

	// Traverse the static array (O(n) time complexity)
	for i := range staticArr {
		fmt.Printf("[%d]:%d, ", i, staticArr[i])
	}

	// Delete an element at index X (O(1) time complexity)
	// Note: In Go, we cannot delete an element from a static array, we use the zero value
	staticArr[1] = 0

	// Print the static array after deletion
	fmt.Println("\nStatic Array after deletion:", staticArr)
}
