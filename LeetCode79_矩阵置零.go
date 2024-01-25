package main

func setZeroes1(matrix [][]int) {
	markLine := map[int]bool{}
	markColumn := map[int]bool{}
	for i, line := range matrix {
		for j, _ := range line {
			if matrix[i][j] == 0 {
				markLine[i] = true
				markColumn[j] = true
			}
		}
	}
	for i := range markLine {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	for i := range markColumn {
		for j := range matrix {
			matrix[j][i] = 0
		}
	}
}
