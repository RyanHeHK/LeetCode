package main

func largest1BorderedSquare(grid [][]int) int {
	maxBorder := 0
	m, n := len(grid), len(grid[0])
	length := make([][]int, m)
	width := make([][]int, m)
	for i := 0; i < m; i++ {
		length[i] = make([]int, n)
		width[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i == 0 && j == 0 {
					width[i][j] = 1
					length[i][j] = 1
				}
				if i == 0 && j > 0 {
					width[i][j] = 1
					length[i][j] = length[i][j-1] + 1
				}
				if j == 0 && i > 0 {
					length[i][j] = 1
					width[i][j] = width[i-1][j] + 1
				}
				if i > 0 && j > 0 {
					width[i][j] = width[i-1][j] + 1
					length[i][j] = length[i][j-1] + 1
				}
			}
			border := min(length[i][j], width[i][j])
			for border > 0 && (length[i-border+1][j] < border || width[i][j-border+1] < border) {
				border--
			}
			maxBorder = max(maxBorder, border)
		}
	}
	return maxBorder * maxBorder
}
