package main

func exist(board [][]byte, word string) bool {

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, word, visited, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, visited [][]bool, i, j, count int) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] {
		return false
	}
	if board[i][j] == word[count] {
		visited[i][j] = true
		if count == len(word)-1 {
			return true
		}
		for x := 0; x < len(directions); x++ {
			if search(board, word, visited, i+directions[x][0], j+directions[x][1], count+1) {
				return true
			}
		}
		visited[i][j] = false
	}

	return false
}
