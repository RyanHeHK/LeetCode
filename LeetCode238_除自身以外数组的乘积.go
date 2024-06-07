package main

func productExceptSelf(nums []int) []int {
	left := []int{}
	right := []int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		left = append(left, 1)
		right = append(right, 1)
	}
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
		right[len(nums)-1-i] = nums[len(nums)-i] * right[len(nums)-i]
	}
	for i := 0; i < len(nums); i++ {
		res = append(res, left[i]*right[i])
	}
	return res
}
