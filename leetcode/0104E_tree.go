package main

func maxDepth(root *TreeNode) int {
	if root == nil { // Base case
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
