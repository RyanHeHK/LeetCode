package main

import "LeetCode/utils"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	revertS := revertString(s)
	for i, b1 := range s {
		for j, b2 := range revertS {
			if b1 == b2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = utils.Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[n][n]
}

func revertString(s string) string {
	b := []byte{}
	for i := range s {
		b = append(b, s[len(s)-i-1])
	}
	revertS := string(b)
	return revertS
}
