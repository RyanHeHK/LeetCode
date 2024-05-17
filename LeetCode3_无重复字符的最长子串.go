package main

import "LeetCode/utils"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := map[int32]int{}
	left := 0
	res := 0
	for i, c := range s {
		if _, ok := m[c]; ok {
			left = utils.Max(left, m[c]+1)
		}
		res = utils.Max(res, i-left+1)
		m[c] = i
	}
	return res
}
