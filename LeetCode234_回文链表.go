package main

func isPalindrome2(head *ListNode) bool {
	node := head
	length := 0
	for node != nil {
		length++
		node = node.Next
	}
	node = head
	for i := 0; i < length/2; i++ {
		node = node.Next
	}
	node = reverseList(node)
	for i := 0; i < length/2; i++ {
		if head.Val != node.Val {
			return false
		}
		head = head.Next
		node = node.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}
