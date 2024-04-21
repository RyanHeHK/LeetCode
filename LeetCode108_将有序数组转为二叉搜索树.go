package main

func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  setNode(nums, 0, mid),
		Right: setNode(nums, mid+1, len(nums)),
	}
	return root
}

func setNode(nums []int, start, end int) *TreeNode {
	if start == end {
		return nil
	}
	mid := (end + start) / 2
	node := &TreeNode{
		Val:   nums[mid],
		Left:  setNode(nums, start, mid),
		Right: setNode(nums, mid+1, end),
	}
	return node
}
