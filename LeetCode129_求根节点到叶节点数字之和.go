package main

import "strconv"

func sumNumbers(root *TreeNode) int {
	vStr := ""
	return sumNodeValue(root, 0, vStr)
}

func sumNodeValue(node *TreeNode, sumV int, vStr string) int {
	if node == nil {
		return sumV
	}
	if node != nil {
		vStr += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			v, _ := strconv.Atoi(vStr)
			sumV += v
			return sumV
		}
		sumV = sumNodeValue(node.Left, sumV, vStr)
		sumV = sumNodeValue(node.Right, sumV, vStr)
	}
	return sumV
}
