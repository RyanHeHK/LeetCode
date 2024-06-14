package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	cur := head
	for i := 1; i < left; i++ {
		pre = pre.Next
		cur = cur.Next
	}
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
