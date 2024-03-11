package main

func maxSubArray(nums []int) int {
	maxValue := make([]int, len(nums))
	maxValue[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+maxValue[i-1] {
			maxValue[i] = nums[i]
		} else {
			maxValue[i] = nums[i] + maxValue[i-1]
		}
	}
	res := maxValue[0]
	for i := 1; i < len(maxValue); i++ {
		if maxValue[i] > res {
			res = maxValue[i]
		}
	}
	return res
}
