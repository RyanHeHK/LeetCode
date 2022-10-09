package main

func generateParenthesis(n int) []string {
	var res []string
	res = formParenthesis(n, n, res, "")
	return res
}

func formParenthesis(left, right int, res []string, curStr string) []string {
	if left == 0 && right == 0 {
		res = append(res, curStr)
		return res
	}

	if left > right {
		return res
	}
	if left > 0 {
		res = formParenthesis(left-1, right, res, curStr+"(")
	}
	if right > 0 {
		res = formParenthesis(left, right-1, res, curStr+")")
	}
	return res
}
