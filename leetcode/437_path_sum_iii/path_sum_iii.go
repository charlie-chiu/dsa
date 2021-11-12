package main

import (
	. "dsa/leetcode"
)

// DFS twice... O(N^2) ?
var result int

func pathSum(root *TreeNode, targetSum int) int {
	result = 0
	preorderTraversal(root, targetSum)

	return result
}

func preorderTraversal(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}

	checkPathSum(root, targetSum)
	preorderTraversal(root.Left, targetSum)
	preorderTraversal(root.Right, targetSum)
}

func checkPathSum(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		result += 1
	}

	checkPathSum(root.Left, targetSum-root.Val)
	checkPathSum(root.Right, targetSum-root.Val)
}
