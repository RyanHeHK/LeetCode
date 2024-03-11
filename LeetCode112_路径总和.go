package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	return nodeValueSum(root, 0, targetSum)
}

func nodeValueSum(node *TreeNode, tmpSum int, targetSum int) bool {
	if node == nil {
		return false
	}
	tmpSum += node.Val
	if tmpSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}
	if nodeValueSum(node.Left, tmpSum, targetSum) {
		return true
	}
	if nodeValueSum(node.Right, tmpSum, targetSum) {
		return true
	}
	return false
}
