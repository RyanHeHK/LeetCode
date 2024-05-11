package main

import "LeetCode/utils"

func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	m := make([][]int, len(matrix))
	for i := range m {
		m[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				m[i][j] = 1
				if i > 0 && j > 0 {
					m[i][j] = utils.Min(utils.Min(m[i-1][j], m[i][j-1]), m[i-1][j-1]) + 1
				}
				maxSide = utils.Max(maxSide, m[i][j])
			}
		}
	}
	return maxSide * maxSide
}
