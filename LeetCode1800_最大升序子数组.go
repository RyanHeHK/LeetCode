package main

import "LeetCode/utils"

func maxAscendingSum(nums []int) int {
	res := nums[0]
	tmp := nums[0]
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > last {
			tmp += nums[i]
		} else {
			res = utils.Max(res, tmp)
			tmp = nums[i]
		}
		last = nums[i]
	}
	res = utils.Max(res, tmp)
	return res
}
