package main

import "LeetCode/utils"

func maxDepth(root *TreeNode) int {
	return calculateLevel(root, 0)
}

func calculateLevel(node *TreeNode, maxDeep int) int {
	if node != nil {
		maxDeep++
		return utils.Max(calculateLevel(node.Left, maxDeep), calculateLevel(node.Right, maxDeep))
	}
	return maxDeep
}
