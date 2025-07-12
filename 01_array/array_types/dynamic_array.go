package array_types

import "fmt"

func DynamicArrayExample() {
	fmt.Println("\nDynamic Array Example:")

	// Dynamic array using a slice in Go
	var dynamicArr []int

	// Append values to the dynamic array (O(1) amortized time complexity)
	dynamicArr = append(dynamicArr, 1, 2, 3, 4, 5)

	// Assign a value at index X (O(1) time complexity)
	dynamicArr[4] = 10

	// Traverse the dynamic array (O(n) time complexity)
	for i := range dynamicArr {
		fmt.Printf("[%d]:%d, ", i, dynamicArr[i])
	}

	// Delete an element at index X (O(n) time complexity)
	// Delete element at index 1
	dynamicArr = append(dynamicArr[:1], dynamicArr[2:]...) // Remove element at index 1

	// Print the dynamic array after deletion
	fmt.Println("\nDynamic Array after deletion:", dynamicArr)
}
