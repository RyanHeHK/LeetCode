package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	node := root
	for node != nil {
		if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Left
		}
	}
	return root
}
