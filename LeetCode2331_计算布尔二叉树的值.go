package main

const (
	Flag_FALSE = 0
	Flag_TRUE  = 1
	Flag_OR    = 2
	Flag_AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	if root.Val == Flag_TRUE {
		return true
	}
	if root.Val == Flag_FALSE {
		return false
	}
	if root.Val == Flag_OR {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	if root.Val == Flag_AND {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
	return false
}
