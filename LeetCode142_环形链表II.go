package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := map[*ListNode]bool{}
	node := head
	for true {
		if m[node] == true {
			break
		} else {
			m[node] = true
			node = node.Next
		}
	}
	return node
}

// 双指针
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
