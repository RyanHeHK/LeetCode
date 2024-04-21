package main

import "LeetCode/utils"

func minSubArrayLen(target int, nums []int) int {
	res := 0
	tmpVal := 0
	tmpLength := 0
	left := 0
	for i, n := range nums {
		tmpVal += n
		tmpLength++
		if target <= tmpVal {
			for left <= i && target <= tmpVal {
				tmpVal -= nums[left]
				if res == 0 {
					res = tmpLength
				} else {
					res = utils.Min(res, tmpLength)
				}
				tmpLength--
				left++
			}
		}
	}
	return res
}
