package main

// https://leetcode.cn/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	index := 0
	res := ""
	standard := strs[0]
	for index < len(strs[0]) {
		c := standard[index]
		for i := 1; i < len(strs); i++ {
			if index > len(strs[i])-1 {
				return res
			}
			if c != strs[i][index] {
				return res
			}
		}
		res += string(c)
		index++
	}
	return res
}
