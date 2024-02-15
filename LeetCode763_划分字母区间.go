package main

func partitionLabels(s string) []int {
	lastIndex := make([]int, 26)
	for i, c := range s {
		lastIndex[c-'a'] = i
	}
	res := []int{}
	start, end := 0, 0
	for i, c := range s {
		if lastIndex[c-'a'] > end {
			end = lastIndex[c-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
