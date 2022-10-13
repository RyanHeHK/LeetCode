package main

func partition(head *ListNode, x int) *ListNode {
	var leftNode, leftTailNode, rightNode, rightTailNode *ListNode
	for head != nil {
		node := &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		if head.Val < x {
			if leftNode == nil {
				leftNode = node
				leftTailNode = node
			} else {
				leftTailNode.Next = node
				leftTailNode = node
			}
		} else {
			if rightNode == nil {
				rightNode = node
				rightTailNode = node
			} else {
				rightTailNode.Next = node
				rightTailNode = node
			}
		}
		head = head.Next
	}
	if leftNode != nil {
		leftTailNode.Next = rightNode
	}
	if leftNode == nil && rightNode != nil {
		leftNode = rightNode
	}
	return leftNode
}
