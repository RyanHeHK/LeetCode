package main

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		if root.Right.Val == root.Left.Val {
			return leftNodeIsSymmetric(root.Left, root.Right)
		} else {
			return false
		}
	} else {
		return false
	}
}

func leftNodeIsSymmetric(left, right *TreeNode) bool {
	in := false
	out := false
	if left.Right == nil && right.Left == nil {
		in = true
	}
	if left.Left == nil && right.Right == nil {
		out = true
	}
	if left.Right != nil && right.Left != nil {
		if left.Right.Val == right.Left.Val {
			in = leftNodeIsSymmetric(left.Right, right.Left)
		}
	}
	if left.Left != nil && right.Right != nil {
		if left.Left.Val == right.Right.Val {
			out = leftNodeIsSymmetric(left.Left, right.Right)
		}
	}
	return in && out
}
