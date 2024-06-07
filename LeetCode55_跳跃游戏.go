package main

import "LeetCode/utils"

// 贪心
func canJump(nums []int) bool {
	rightMost := 0
	for i := range nums {
		if i <= rightMost {
			rightMost = utils.Max(rightMost, i+nums[i])
			if rightMost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
