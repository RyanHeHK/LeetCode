package main

// https://leetcode.cn/problems/excel-sheet-column-title/
func convertToTitle(columnNumber int) string {
	m := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	curStr := ""
	for columnNumber > 0 {
		columnNumber--
		yu := columnNumber % 26
		curStr = m[yu] + curStr
		columnNumber = columnNumber / 26
	}
	return curStr
}
