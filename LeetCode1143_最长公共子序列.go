package main

import "LeetCode/utils"

func longestCommonSubsequence(text1 string, text2 string) int {
	m := make([][]int, len(text1)+1)
	for i := range m {
		m[i] = make([]int, len(text2)+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				m[i+1][j+1] = m[i][j] + 1
			} else {
				m[i+1][j+1] = utils.Max(m[i+1][j], m[i][j+1])
			}
		}
	}
	return m[len(text1)][len(text2)]
}
