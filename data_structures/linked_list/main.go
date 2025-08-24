package main

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/doubly"
	"github.com/WolfieLeader/data-structures-algorithms/data_structures/linked_list/singly"
)

func main() {
	fmt.Println("Singly Linked List Example:")
	singlyExample()
	fmt.Println()

	fmt.Println("Doubly Linked List Example:")
	doublyExample()
	fmt.Println()
}

func singlyExample() {
	s := singly.New[int]()
	fmt.Printf("- Empty:    %v\n", s)
	s.AddLast(1)
	fmt.Printf("- Element:  %v\n", s)
	s.AddFirst(5, 4, 3, 2)
	fmt.Printf("- More:     %v\n", s)
}

func doublyExample() {
	d := doubly.New[int]()
	fmt.Printf("- Empty:    %v\n", d)
	d.AddLast(1)
	fmt.Printf("- Element:  %v\n", d)
	d.AddFirst(5, 4, 3, 2)
	fmt.Printf("- More:     %v\n", d)
}
