package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	bst "github.com/WolfieLeader/data-structures-algorithms/data_structures/tree/binary_search_tree"
)

func main() {
	fmt.Println("Binary Search Tree Example")
	bstExample()
	fmt.Println()
}

func bstExample() {
	t := bst.New[int]()
	t.InsertI(10, 5, 15, 2, 5)
	t.InsertR(7, 13, 17, 10)
	fmt.Printf("- Tree:\n%v", t)

	root, _ := t.Root()
	minI, _ := t.MinI()
	maxI, _ := t.MaxI()
	minR, _ := t.MinR()
	maxR, _ := t.MaxR()
	fmt.Printf("- Size: %d, Root: %v, Height: %d(ok=%t)\n", t.Size(), root, t.HeightI(), t.HeightR() == t.HeightI())
	fmt.Printf("- Min: %v(ok=%t), Max: %v(ok=%t), Is Balanced: %t\n", minI, minI == minR, maxI, maxI == maxR, t.IsBalanced())
	fmt.Printf("- Contains 17: %t(ok=%t)\n", t.ContainsI(17), t.ContainsI(17) == t.ContainsR(17))
	fmt.Printf("- Contains 0: %t(ok=%t)\n", t.ContainsI(0), t.ContainsI(0) == t.ContainsR(0))

	dfsInOrderI, dfsInOrderR := traverse(t.TraverseInOrderI), traverse(t.TraverseInOrderR)
	dfsPreOrderI, dfsPreOrderR := traverse(t.TraversePreOrderI), traverse(t.TraversePreOrderR)
	dfsPostOrderI, dfsPostOrderR := traverse(t.TraversePostOrderI), traverse(t.TraversePostOrderR)
	bfs := traverse(t.TraverseBreadthFirst)
	fmt.Printf("- DFS In Order: %s(ok=%t)\n", dfsInOrderI, dfsInOrderI == dfsInOrderR)
	fmt.Printf("- DFS Pre Order: %s(ok=%t)\n", dfsPreOrderI, dfsPreOrderI == dfsPreOrderR)
	fmt.Printf("- DFS Post Order: %s(ok=%t)\n", dfsPostOrderI, dfsPostOrderI == dfsPostOrderR)
	fmt.Printf("- BFS: %s\n", bfs)

	for _, v := range []int{6, 10, 13, 100, -1} {
		pi, piOk := t.PredecessorI(v)
		pr, prOk := t.PredecessorR(v)
		if piOk && prOk {
			fmt.Printf("- %d Predecessor: %v(ok=%t), ", v, pi, pi == pr)
		} else {
			fmt.Printf("- %d Predecessor: none, ", v)
		}

		si, siOk := t.SuccessorI(v)
		sr, srOk := t.SuccessorR(v)
		if siOk && srOk {
			fmt.Printf("%d Successor: %v(ok=%t)\n", v, si, si == sr)
		} else {
			fmt.Printf("%d Successor: none\n", v)
		}
	}

	fmt.Printf("- Delete 17, 5, 10, 2: %v\n", t.DeleteI(17, 5) == t.DeleteR(10, 2))
	fmt.Printf("- After deletes:\n%v", t)
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", t.Size(), t.HeightI(), t.HeightR() == t.HeightI(), t.IsBalanced())
	t.Clear()
	fmt.Printf("- After clear:\n%v", t)
	fmt.Printf("- Size: %d, IsEmpty: %t, Is Balanced: %t\n", t.Size(), t.IsEmpty(), t.IsBalanced())

	t.InsertI(1, 2, 3)
	t.InsertR(4, 5, 6)
	fmt.Printf("- Skewed tree:\n%v", t)
	fmt.Printf("- To Slice: %v(ok=%t)\n", t.ToSliceI(), slices.Equal(t.ToSliceI(), t.ToSliceR()))
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", t.Size(), t.HeightI(), t.HeightR() == t.HeightI(), t.IsBalanced())
	c1 := t.CopyI()
	c1.InsertI(0)
	c2 := t.CopyR()
	c2.InsertR(7)
	fmt.Printf("- Are BSTs equal? (skew tree + 0) == (skew tree + 7): %t(ok=%t)\n", c1.EqualI(c2), c1.EqualI(c2) == c1.EqualR(c2))
	c1.DeleteI(0)
	c2.DeleteR(7)
	fmt.Printf("- Are BSTs equal? (skew tree) == (skew tree + 0 - 0) == (skew tree + 7 - 7): %t(ok=%t)\n", t.EqualI(c1) && t.EqualI(c2), t.EqualI(c1) && t.EqualI(c2) == t.EqualR(c1) && t.EqualR(c2))
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
