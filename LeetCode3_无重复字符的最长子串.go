package main

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := make(map[int32]int)
	left := 0
	max := 0
	for i, c := range s {
		if _, ok := m[c]; ok {
			if left < m[c]+1 {
				left = m[c] + 1
			}
		}
		m[c] = i
		if max < i+1-left {
			max = i + 1 - left
		}
	}
	return max
}
