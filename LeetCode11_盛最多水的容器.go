package main

// https://leetcode.cn/problems/container-with-most-water/
// 双指针
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxSize := 0
	for left < right {
		if height[left] > height[right] {
			if maxSize < (right-left)*height[right] {
				maxSize = (right - left) * height[right]
			}
			right--
		} else {
			if maxSize < (right-left)*height[left] {
				maxSize = (right - left) * height[left]
			}
			left++
		}
	}
	return maxSize
}
