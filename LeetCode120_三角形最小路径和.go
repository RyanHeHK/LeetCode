package main

import (
	"LeetCode/utils"
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	res := make([]int, n)
	res[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			res[j] = utils.Min(res[j-1], res[j]) + triangle[i][j]
		}
		res[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = utils.Min(ans, res[i])
	}
	return ans
}
