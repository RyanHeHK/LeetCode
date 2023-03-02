package main

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i < n/2; i++ {
		for j := 0; j < len(matrix)-1; j++ {
			tmp := matrix[i][j]
			matrix[n-j][i] = matrix[n-i][0]
			matrix[n-i][0] = matrix[n][n-i]
			matrix[n][n-i] = matrix[j][n]
			matrix[j][n] = tmp
		}
	}
}
