package main

func reverseList024(head *ListNode) *ListNode {
	var pre *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
