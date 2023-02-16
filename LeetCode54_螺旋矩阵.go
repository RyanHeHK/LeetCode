package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row := len(matrix)
	column := len(matrix[0])
	res := []int{}
	total := row * column
	i, j := 0, 0
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dindex := 0
	visited := make([][]bool, row)
	for k := 0; k < row; k++ {
		visited[k] = make([]bool, column)
	}
	for total > 0 {
		res = append(res, matrix[i][j])
		visited[i][j] = true
		i += direct[dindex][0]
		j += direct[dindex][1]
		if i >= row || j >= column || i < 0 || j < 0 || visited[i][j] {
			i -= direct[dindex][0]
			j -= direct[dindex][1]
			dindex = (dindex + 1) % 4
			i += direct[dindex][0]
			j += direct[dindex][1]
		}
		total--
	}
	return res
}
