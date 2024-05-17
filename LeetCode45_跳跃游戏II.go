package main

import "LeetCode/utils"

func jump(nums []int) int {
	length := len(nums)
	step := 0
	end := 0
	maxLength := 0
	for i := 0; i < length-1; i++ {
		maxLength = utils.Max(maxLength, nums[i]+i)
		if maxLength >= length-1 {
			step++
			break
		}
		if i == end {
			end = maxLength
			step++
		}
	}
	return step
}
