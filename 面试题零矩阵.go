package main

func setZeroes(matrix [][]int) {
	m := len(matrix[0])
	n := len(matrix)
	points := [][]int{}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				points = append(points, []int{i, j})
			}
		}
	}
	for _, point := range points {
		for i := 0; i < m; i++ {
			matrix[point[0]][i] = 0
		}
		for i := 0; i < n; i++ {
			matrix[i][point[1]] = 0
		}
	}
}
