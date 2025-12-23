package main

// BFS used

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// HACK: IMPORTANT! Add Root
	queue := []*TreeNode{root}

	out := make([][]int, 0)

	for len(queue) > 0 {
		// HACK: IMPORTANT! Get Size at first
		levelSize := len(queue)
		inner := make([]int, 0, levelSize)

		// NOTE: LOOP over each item in the same level
		for range levelSize {
			node := queue[0]  // Get first node
			queue = queue[1:] // Pop
			inner = append(inner, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		out = append(out, inner)
	}

	return out
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
