package main

func generateMatrix(n int) [][]int {
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	dindex := 0
	total := n * n
	val := 1
	i, j := 0, 0
	for total > 0 {
		res[i][j] = val
		val++
		i += direct[dindex][0]
		j += direct[dindex][1]
		if i == n-1 || j == n-1 || i == 0 || j == 0 || res[i][j] != 0 {
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
