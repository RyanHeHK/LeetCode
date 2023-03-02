package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else {
		if p.Val == q.Val {
			if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
