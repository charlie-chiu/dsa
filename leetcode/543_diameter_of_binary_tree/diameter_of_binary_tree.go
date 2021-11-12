package main

import . "dsa/leetcode"

// first approach, 0ms 100%, 4.4MB, 93.32%

var globalMax int

func diameterOfBinaryTree(root *TreeNode) int {
	globalMax = 0
	helper(root)

	return globalMax
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left)
	r := helper(root.Right)
	globalMax = max(globalMax, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
