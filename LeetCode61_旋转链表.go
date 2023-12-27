package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	node := head
	length := 0
	tail := &ListNode{}
	for node != nil {
		length += 1
		tail = node
		node = node.Next
	}
	if k%length == 0 {
		return head
	} else {
		node = head
		for i := 0; i < length-(k%length)-1; i++ {
			node = node.Next
		}
		tail.Next = head
		res := node.Next
		node.Next = nil
		return res
	}
}
