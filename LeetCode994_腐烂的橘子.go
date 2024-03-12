package main

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	if !hasFreshOrange(grid) {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	time := getTime(grid)

	if hasFreshOrange(grid) {
		return -1
	}
	return time
}

func hasFreshOrange(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return true
			}
		}
	}
	return false
}

func getTime(grid [][]int) int {
	q := make([][]int, 0)
	for i := range grid {
		for j, val := range grid[i] {
			if 2 == val {
				q = append(q, []int{i, j})
			}
		}
	}
	time := 0
	for len(q) > 0 {
		next := [][]int{}
		for i := 0; i < len(q); i++ {
			r := q[i][0]
			l := q[i][1]
			next = getNext(grid, r+1, l, next)
			next = getNext(grid, r-1, l, next)
			next = getNext(grid, r, l+1, next)
			next = getNext(grid, r, l-1, next)
		}
		if len(next) != 0 {
			time++
		}
		q = next
	}
	return time
}

func getNext(grid [][]int, i, j int, queue [][]int) [][]int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return queue
	}
	if grid[i][j] == 1 {
		grid[i][j] = 2
		queue = append(queue, []int{i, j})
	}
	return queue
}
