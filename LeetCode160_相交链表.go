package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1 := headA
	node2 := headB
	for node1 != node2 {
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return node1
}
