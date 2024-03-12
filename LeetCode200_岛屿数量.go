package main

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfsIslands(grid, i, j)
			}
		}
	}
	return res
}

func dfsIslands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfsIslands(grid, i+1, j)
		dfsIslands(grid, i-1, j)
		dfsIslands(grid, i, j+1)
		dfsIslands(grid, i, j-1)
	}
}
