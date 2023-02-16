package main

func exist(board [][]byte, word string) bool {

	visted := make([][]bool, len(board))
	for i := range visted {
		visted[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if len(word) == 1 {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, word string, visited [][]int, i, j int) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visted[i][j] = true
	for {
		if search {

		}
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	return false
}
