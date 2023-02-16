package main

import "sort"

// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				arr := []int{nums[i], nums[L], nums[R]}
				res = append(res, arr)
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			}
			if sum < 0 {
				L++
			}
			if sum > 0 {
				R--
			}
		}
	}
	return res
}
