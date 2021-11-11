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

// from leetcode 101
func isBalanced2(root *TreeNode) bool {
	return root == nil || helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	maxDepthLeft := helper(root.Left)
	if maxDepthLeft == -1 {
		return -1
	}
	maxDepthRight := helper(root.Right)
	if maxDepthRight == -1 {
		return -1
	}

	if abs(maxDepthLeft, maxDepthRight) > 1 {
		return -1
	} else {
		return max(maxDepthLeft, maxDepthRight) + 1
	}
}
