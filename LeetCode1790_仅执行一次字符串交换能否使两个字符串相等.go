package main

// https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
func areAlmostEqual(s1 string, s2 string) bool {
	count := 0
	index := -1
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			if count == 2 {
				return false
			}
			if index != -1 {
				if s1[i] != s2[index] || s1[index] != s2[i] {
					return false
				}
			}
			index = i
			count++
		}
	}
	if count == 1 {
		return false
	}
	return true
}
