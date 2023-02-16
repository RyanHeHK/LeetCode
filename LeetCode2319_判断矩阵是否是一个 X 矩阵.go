package main

// https://leetcode.cn/problems/check-if-matrix-is-x-matrix/
func checkXMatrix(grid [][]int) bool {
	res := true
	l := len(grid)
	for i, v1 := range grid {
		for j, v2 := range v1 {
			if i == j || j == l-i-1 {
				if v2 == 0 {
					return false
				}
			} else {
				if v2 != 0 {
					return false
				}
			}
		}
	}
	return res
}
