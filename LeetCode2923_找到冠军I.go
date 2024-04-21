package main

func findChampion(grid [][]int) int {
	for i := range grid {
		sum := 0
		for j := range grid[0] {
			if grid[i][j] == 1 {
				sum++
			}
		}
		if sum == len(grid)-1 {
			return i
		}
	}
	return -1
}
