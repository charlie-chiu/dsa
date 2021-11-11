package main

import (
	"math"

	. "dsa/leetcode"
)

// brute force, got 4ms, 97.87%, 5.9MB 100% on leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(maxDepth(root.Left), maxDepth(root.Right)) <= 1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
