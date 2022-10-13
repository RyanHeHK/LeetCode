package main

func numComponents(head *ListNode, nums []int) int {
	res := 0
	m := make(map[int]bool)
	exist := false
	for _, v := range nums {
		m[v] = true
	}
	for head != nil {
		if m[head.Val] {
			exist = true
		}
		if exist && !m[head.Val] {
			exist = false
			res++
		}
		head = head.Next
	}
	if exist {
		res++
	}
	return res
}
