package main

// BFS used

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	pushFn := func(node *TreeNode) { queue = append(queue, node) }
	popFn := func() *TreeNode {
		node := queue[0]
		queue[0] = nil
		queue = queue[1:]
		return node
	}

	out := make([][]int, 0)

	for len(queue) > 0 {
		// HACK: IMPORTANT!
		levelSize := len(queue)
		inner := make([]int, 0, levelSize)

		// NOTE: LOOP over each item in the same level
		for range levelSize {
			node := popFn()
			inner = append(inner, node.Val)

			if node.Left != nil {
				pushFn(node.Left)
			}
			if node.Right != nil {
				pushFn(node.Right)
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
