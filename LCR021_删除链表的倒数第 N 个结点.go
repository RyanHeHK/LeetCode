package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := head
	for i := 0; i < n; i++ {
		second = second.Next
	}
	for second != nil {
		first = first.Next
		second = second.Next
	}
	if first.Next == nil {
		dummy.Next = dummy.Next.Next
		return dummy
	}
	first.Next = first.Next.Next
	return dummy.Next
}
