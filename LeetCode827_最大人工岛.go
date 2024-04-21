package main

import "LeetCode/utils"

func largestIsland(grid [][]int) (ans int) {
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	tag := make([][]int, n)
	for i, _ := range tag {
		tag[i] = make([]int, n)
	}
	markIndex := 1
	area := make(map[int]int)
	var markLand func(int, int)
	markLand = func(i, j int) {
		if i < 0 || i > n-1 || j < 0 || j > n-1 {
			return
		}
		if tag[i][j] == 0 && grid[i][j] != 0 {
			tag[i][j] = markIndex
			area[markIndex]++
		} else {
			return
		}
		for _, d := range dir4 {
			markLand(i+d.x, j+d.y)
		}
	}

	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] != 0 {
				markLand(i, j)
				ans = utils.Max(ans, area[markIndex])
				markIndex++
			}
		}
	}

	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == 0 {
				conn := make(map[int]bool)
				conArea := 1
				for _, d := range dir4 {
					if i+d.x >= 0 && i+d.x < n && j+d.y >= 0 && j+d.y < n {
						if tag[i+d.x][j+d.y] != 0 && conn[tag[i+d.x][j+d.y]] == false {
							conArea += area[tag[i+d.x][j+d.y]]
							conn[tag[i+d.x][j+d.y]] = true
						}
					}
				}
				ans = utils.Max(ans, conArea)
			}
		}
	}
	return ans
}
