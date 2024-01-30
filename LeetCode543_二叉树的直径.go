package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxDeep := 0
	if root == nil {
		return maxDeep
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L := dfs(node.Left)
		R := dfs(node.Right)
		maxDeep = max(maxDeep, L+R)
		return 1 + max(L, R)
	}
	dfs(root)
	return maxDeep

}
