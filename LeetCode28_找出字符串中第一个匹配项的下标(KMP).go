package main

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	// 预处理next数组
	next := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	// 在文本串中查找模式串
	j = 0
	for i := 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
