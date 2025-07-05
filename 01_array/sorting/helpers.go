package sorting

import "fmt"

func printName(name string, array []int) { fmt.Printf("\n%s:\nInitial array: %v\n", name, array) }
func printIteration(i int, array []int)  { fmt.Printf("- Iteration %d: %v\n", i+1, array) }
func printFinal(array []int)             { fmt.Printf("Sorted array: %v\n", array) }
