package helpers

import "fmt"

func PrintName(name string, array []int) { fmt.Printf("\n%s:\nInitial array: %v\n", name, array) }
func PrintIteration(i int, array []int)  { fmt.Printf("- Iteration %d: %v\n", i+1, array) }
func PrintFinal(array []int)             { fmt.Printf("Sorted array: %v\n", array) }
