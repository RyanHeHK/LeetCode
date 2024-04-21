package main

import "LeetCode/utils"

func maxProduct(nums []int) int {
	maxRes := nums[0]
	curMax := nums[0]
	curMin := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := curMax
			curMax = curMin
			curMin = temp
		}
		curMax = utils.Max(curMax*nums[i], nums[i])
		curMin = utils.Min(curMin*nums[i], nums[i])
		maxRes = utils.Max(curMax, maxRes)
	}
	return maxRes
}
