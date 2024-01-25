package main

func alternatingSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			length := j - i + 1
			if nums[j]-nums[i] == (length-1)%2 {
				res = max(length, res)
			} else {
				break
			}
		}
	}
	return res
}
