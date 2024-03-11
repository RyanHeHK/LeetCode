package main

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
		curMax = max(curMax*nums[i], nums[i])
		curMin = min(curMin*nums[i], nums[i])
		maxRes = max(curMax, maxRes)
	}
	return maxRes
}
