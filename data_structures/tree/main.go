package main

import (
	"fmt"

	bst "github.com/WolfieLeader/data-structures-algorithms/data_structures/tree/binary_search_tree"
)

func main() {
	bst := bst.New[int]()
	bst.InsertI(5)
	bst.InsertR(3)
	bst.InsertI(7)
	bst.InsertR(2)
	bst.InsertI(15)
	bst.InsertR(33)
	bst.InsertI(72)
	bst.InsertR(24)
	fmt.Println(bst)
}
