package main

import (
	"fmt"

	hashmap "github.com/WolfieLeader/data-structures-algorithms/data_structures/hash_map/map"
	hashset "github.com/WolfieLeader/data-structures-algorithms/data_structures/hash_map/set"
)

func main() {
	fmt.Println("HashMap Example")
	mapExample()
	fmt.Println()

	fmt.Println("HashSet Example")
	setExample()
	fmt.Println()
}

func mapExample() {
	m := hashmap.New[string, int]()
	m.Set("john", 25)
	m.Set("jane", 30)
	m.Set("doe", 22)
	m.Set("john", 28)
	fmt.Printf("- Map: %v\n", m)
	m.Delete("doe")
	fmt.Printf("- Map after deletion: %v\n", m)
	fmt.Printf("- Keys: %v\n", m.Keys())
	fmt.Printf("- Values: %v\n", m.Values())
}

func setExample() {
	m := hashset.New[int]()
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(1)
	fmt.Printf("- Set: %v\n", m)
	m.Delete(2)
	fmt.Printf("- Set after removal: %v\n", m)
	fmt.Printf("- Contains 1: %v\n", m.Contains(1))
	fmt.Printf("- Contains 2: %v\n", m.Contains(2))
}
