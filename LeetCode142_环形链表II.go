package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for true {
		if fast.Next == nil {
			return nil
		}
		if fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			fast = head
			break
		}
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
