package main

func nextLargerNodes(head *ListNode) []int {
	cur := head
	var ans []int
	stack := []int{}
	var dfs func(cur *ListNode, count int)
	dfs = func(cur *ListNode, count int) {
		if cur == nil {
			ans = make([]int, count)
			return
		}
		dfs(cur.Next, count+1)
		for len(stack) > 0 && cur.Val >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && cur.Val < stack[len(stack)-1] {
			ans[count] = stack[len(stack)-1]
		}
		stack = append(stack, cur.Val)
	}
	dfs(cur, 0)
	return ans
}
