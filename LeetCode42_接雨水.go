package main

import "LeetCode/utils"

// 双指针
func trap(height []int) int {
	l := height[0]
	r := height[len(height)-1]
	heightLeft := make([]int, len(height))
	heightRight := make([]int, len(height))
	res := 0
	for i := range height {
		l = utils.Max(l, height[i])
		heightLeft[i] = l
		r = utils.Max(r, height[len(height)-1-i])
		heightRight[len(height)-1-i] = r
	}
	for i := range heightLeft {
		if heightLeft[i] < heightRight[i] {
			res += heightLeft[i] - height[i]
		} else {
			res += heightRight[i] - height[i]
		}
	}
	return res
}
