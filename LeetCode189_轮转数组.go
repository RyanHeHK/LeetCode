package main

func rotateArray(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	res := []int{}
	res = append(res, nums[length-k:]...)
	res = append(res, nums[0:length-k]...)
	copy(nums, res)
}
