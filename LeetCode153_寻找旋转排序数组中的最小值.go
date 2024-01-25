package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := (right + left) / 2
	for left < right {
		if nums[left] <= nums[mid] && nums[mid] <= nums[right] {
			break
		} else if nums[mid] >= nums[left] && nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (right + left) / 2
	}
	return nums[left]
}
