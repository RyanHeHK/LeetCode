package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	layer := []*TreeNode{root}
	if root != nil {
		res = executeSameLayerNode(layer, res, true)
	}
	return res
}

func executeSameLayerNode(layer []*TreeNode, res [][]int, isRight bool) [][]int {
	if len(layer) == 0 {
		return res
	}
	nextLayer := []*TreeNode{}
	layerValue := []int{}
	for _, node := range layer {
		if isRight {
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
		} else {
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		layerValue = append(layerValue, node.Val)
	}
	res = append(res, layerValue)
	res = executeSameLayerNode(nextLayer, res, !isRight)
	return res
}
