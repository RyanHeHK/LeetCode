package main

func sortList(head *ListNode) *ListNode {
	return sortList148(head, nil)
}

func sortList148(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow := head
	fast := head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return merge148(sortList148(head, slow), sortList148(slow, tail))
}

func merge148(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tmp.Next = head1
			head1 = head1.Next
		} else {
			tmp.Next = head2
			head2 = head2.Next
		}
		tmp = tmp.Next
	}
	for head1 != nil {
		tmp.Next = head1
		head1 = head1.Next
		tmp = tmp.Next
	}
	for head2 != nil {
		tmp.Next = head2
		head2 = head2.Next
		tmp = tmp.Next
	}
	return dummyHead.Next
}
