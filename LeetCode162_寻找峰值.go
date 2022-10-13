package main

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
	//if len(nums) == 1 {
	//	return 0
	//}
	//for i := 0; i < len(nums); i++ {
	//	if i-1 >= 0 && i+1 < len(nums) {
	//		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
	//			return i
	//		}
	//	} else {
	//		if i-1 < 0 {
	//			if nums[i] > nums[i+1] {
	//				return i
	//			}
	//		}
	//		if i+1 >= len(nums) {
	//			if nums[i] > nums[i-1] {
	//				return i
	//			}
	//		}
	//	}
	//}
	//return -1
}
