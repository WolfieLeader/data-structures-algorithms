package main

import (
	"cmp"
	"fmt"
	"strings"

	bst "github.com/WolfieLeader/data-structures-algorithms/data_structures/tree/binary_search_tree"
)

func main() {
	fmt.Println("Binary Search Tree Example")
	bstExample()
	fmt.Println()
}

func bstExample() {
	tree := bst.New[int]()
	tree.InsertI(10, 5, 15, 2, 5)
	tree.InsertR(7, 13, 17, 10)
	fmt.Printf("- Tree:\n%v", tree)

	root, _ := tree.Root()
	minI, _ := tree.MinI()
	maxI, _ := tree.MaxI()
	minR, _ := tree.MinR()
	maxR, _ := tree.MaxR()
	fmt.Printf("- Size: %d, Root: %v, Height: %d(ok=%t)\n", tree.Size(), root, tree.HeightI(), tree.HeightR() == tree.HeightI())
	fmt.Printf("- Min: %v(ok=%t), Max: %v(ok=%t), Is Balanced: %t\n", minI, minI == minR, maxI, maxI == maxR, tree.IsBalanced())
	fmt.Printf("- Contains 17: %t(ok=%t)\n", tree.ContainsI(17), tree.ContainsI(17) == tree.ContainsR(17))
	fmt.Printf("- Contains 0: %t(ok=%t)\n", tree.ContainsI(0), tree.ContainsI(0) == tree.ContainsR(0))

	dfsInOrderI, dfsInOrderR := traverse(tree.TraverseInOrderI), traverse(tree.TraverseInOrderR)
	dfsPreOrderI, dfsPreOrderR := traverse(tree.TraversePreOrderI), traverse(tree.TraversePreOrderR)
	dfsPostOrderI, dfsPostOrderR := traverse(tree.TraversePostOrderI), traverse(tree.TraversePostOrderR)
	bfs := traverse(tree.TraverseBreadthFirst)
	fmt.Printf("- DFS In Order: %s(ok=%t)\n", dfsInOrderI, dfsInOrderI == dfsInOrderR)
	fmt.Printf("- DFS Pre Order: %s(ok=%t)\n", dfsPreOrderI, dfsPreOrderI == dfsPreOrderR)
	fmt.Printf("- DFS Post Order: %s(ok=%t)\n", dfsPostOrderI, dfsPostOrderI == dfsPostOrderR)
	fmt.Printf("- BFS: %s\n", bfs)

	for _, v := range []int{6, 10, 13, 20, 100, -1} {
		predI, predIOk := tree.PredecessorI(v)
		predR, predROk := tree.PredecessorR(v)
		succI, succIOk := tree.SuccessorI(v)
		succR, succROk := tree.SuccessorR(v)
		if predIOk && predROk {
			fmt.Printf("- %d Predecessor: %v(ok=%t), ", v, predI, predI == predR)
		} else {
			fmt.Printf("- %d Predecessor: none, ", v)
		}
		if succIOk && succROk {
			fmt.Printf("%d Successor: %v(ok=%t)\n", v, succI, succI == succR)
		} else {
			fmt.Printf("%d Successor: none\n", v)
		}
	}

	fmt.Printf("- Delete 17, 5, 10, 2: %v\n", tree.DeleteI(17) && tree.DeleteR(5) && tree.DeleteI(10) && tree.DeleteR(2))
	fmt.Printf("- After deletes:\n%v", tree)
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", tree.Size(), tree.HeightI(), tree.HeightR() == tree.HeightI(), tree.IsBalanced())
	tree.Clear()
	fmt.Printf("- After clear:\n%v", tree)
	fmt.Printf("- Size: %d, IsEmpty: %t, Is Balanced: %t\n", tree.Size(), tree.IsEmpty(), tree.IsBalanced())

	tree.InsertI(1, 2, 3)
	tree.InsertR(4, 5, 6)
	fmt.Printf("- Skewed tree:\n%v", tree)
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", tree.Size(), tree.HeightI(), tree.HeightR() == tree.HeightI(), tree.IsBalanced())
}

func traverse[T cmp.Ordered](method func(fn func(value T))) string {
	var sb strings.Builder
	first := true
	method(func(v T) {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", v)
		first = false
	})
	return sb.String()
}
