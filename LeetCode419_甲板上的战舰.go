package main

func countBattleships(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				searchShip(board, i, j)
				count++
			}
		}
	}
	return count
}

func searchShip(board [][]byte, i, j int) {
	if board[i][j] == 'X' {
		board[i][j] = '.'
	} else {
		return
	}
	searchShip(board, i+1, j)
	searchShip(board, i, j+1)
}
