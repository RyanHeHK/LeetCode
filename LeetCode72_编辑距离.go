package main

import "LeetCode/utils"

func minDistance(word1 string, word2 string) int {
	m := make([][]int, len(word1)+1)
	for i := range m {
		m[i] = make([]int, len(word2)+1)
		if i == 0 {
			for j := range m[i] {
				m[i][j] = j
			}
		}
		m[i][0] = i
	}

	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				m[i+1][j+1] = m[i][j]
			} else {
				m[i+1][j+1] = utils.Min(utils.Min(m[i][j], m[i+1][j]), m[i][j+1]) + 1
			}
		}
	}
	return m[len(word1)][len(word2)]
}
