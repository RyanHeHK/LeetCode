package main

import "strconv"

// https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
func isPalindrome(x int) bool {
	res := true
	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for true {
		if left > right {
			break
		}
		if s[left] != s[right] {
			res = false
			return res
		}
		left++
		right--
	}
	return res
}
