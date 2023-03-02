package main

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	levelNode := [][]*TreeNode{{root}}
	for i := 0; i < len(levelNode); i++ {
		levelVal := []int{}
		levelNodeTmp := []*TreeNode{}
		for _, node := range levelNode[i] {
			levelVal = append(levelVal, node.Val)
			if node.Left != nil {
				levelNodeTmp = append(levelNodeTmp, node.Left)
			}
			if node.Right != nil {
				levelNodeTmp = append(levelNodeTmp, node.Right)
			}
			levelNode = append(levelNode, levelNodeTmp)
		}
		ret = append(ret, levelVal)
	}
	return ret
}
