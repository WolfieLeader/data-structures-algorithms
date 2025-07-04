package arrays

import "fmt"

func MultiDimensionalArrayExample() {
	fmt.Println("\nMulti-Dimensional Array Example:")

	var matrixArr [3][3]string

	// Initialize the multi-dimensional array
	matrixArr = [3][3]string{
		{"O", "X", "O"},
		{" ", " ", "X"},
		{"X", " ", "O"},
	}

	// Traverse the multi-dimensional array (O(n) time complexity)
	for i := range matrixArr {
		for j := range matrixArr[i] {
			fmt.Printf("%s ", matrixArr[i][j])
		}
		fmt.Println() // New line after each row
	}

	// Assign a value at index (O(1) time complexity)
	matrixArr[1][2] = "O"

	// Print the multi-dimensional array after assignment
	fmt.Println("\nMulti-Dimensional Array after assignment:", matrixArr)
}
