package main

type RandomCopyNode struct {
	Val    int
	Next   *RandomCopyNode
	Random *RandomCopyNode
}

func copyRandomList(head *RandomCopyNode) *RandomCopyNode {
	if head == nil {
		return nil
	}

	clone := make(map[*RandomCopyNode]*RandomCopyNode)

	// 1) Create clone nodes
	curr := head
	for curr != nil {
		clone[curr] = &RandomCopyNode{Val: curr.Val}

		curr = curr.Next
	}

	// 2) Connect Next and Random
	curr = head
	for curr != nil {
		clonedNode := clone[curr]

		clonedNode.Next = clone[curr.Next]     // nil maps to nil (ok)
		clonedNode.Random = clone[curr.Random] // nil maps to nil (ok)

		curr = curr.Next
	}

	return clone[head]
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
