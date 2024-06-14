package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	// 添加一个哑节点作为新链表的头部
	dummy := &ListNode{Next: head}
	prev := dummy

	// 找到反转区间的前一个节点
	for i := 1; i < m; i++ {
		prev = prev.Next
	}

	// 反转区间的第一个节点
	curr := prev.Next

	// 反转区间内的节点
	for i := m; i < n; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
