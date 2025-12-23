package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS used

func invertTree(root *TreeNode) *TreeNode {
	if root == nil { // Base condition
		return nil
	}

	// swap
	root.Left, root.Right = root.Right, root.Left

	// recurse
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
