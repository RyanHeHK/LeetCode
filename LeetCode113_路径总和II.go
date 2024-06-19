package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	path := []int{}
	dfs(root, targetSum, 0, path, &ans)
	return ans
}

func dfs(node *TreeNode, targetSum int, currentSum int, path []int, ans *[][]int) {
	if node == nil {
		return
	}

	currentSum += node.Val
	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil && currentSum == targetSum {
		// 将符合条件的路径加入结果中
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
	}

	dfs(node.Left, targetSum, currentSum, path, ans)
	dfs(node.Right, targetSum, currentSum, path, ans)

	path = path[:len(path)-1]
}
