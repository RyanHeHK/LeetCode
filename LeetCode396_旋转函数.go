package main

import "LeetCode/utils"

// https://leetcode.cn/problems/rotate-function/
func maxRotateFunction(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	f := 0
	for i, num := range nums {
		f += i * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = utils.Max(ans, f)
	}
	return ans
}
