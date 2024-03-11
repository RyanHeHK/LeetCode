package main

func sortColors(nums []int) {
	leftIndex := 0
	rightIndex := len(nums) - 1
	for i := 0; i <= rightIndex; i++ {
		for ; i <= rightIndex && nums[i] == 2; rightIndex-- {
			nums[i], nums[rightIndex] = nums[rightIndex], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[leftIndex] = nums[leftIndex], nums[i]
			leftIndex++
		}
	}
}
