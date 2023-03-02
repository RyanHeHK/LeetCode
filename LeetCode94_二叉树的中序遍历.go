package main

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	middleScan(root, &res)
	return res
}

func middleScan(node *TreeNode, res *[]int) *[]int {
	if node != nil {
		middleScan(node.Left, res)
		*res = append(*res, node.Val)
		middleScan(node.Right, res)
	}
	return res
}
