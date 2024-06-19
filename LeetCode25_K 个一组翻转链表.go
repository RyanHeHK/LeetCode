package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	cur := head
	pre := dummy
	count := 1
	for cur != nil {
		cur = cur.Next
		count++
		if count == k && cur != nil {
			cur = pre.Next
			for i := 1; i < k; i++ {
				next := cur.Next
				cur.Next = next.Next
				next.Next = pre.Next
				pre.Next = next
			}
			pre = cur
			count = 0
		}
	}
	return dummy.Next
}
