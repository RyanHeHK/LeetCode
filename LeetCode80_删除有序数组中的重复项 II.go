package main

func removeDuplicates(nums []int) int {
	index := 0
	same := false
	for i := range nums {
		if i == 0 {
			index++
			continue
		}
		if nums[i] == nums[index-1] && !same {
			nums[index] = nums[i]
			same = true
			index++
			continue
		}
		if nums[i] == nums[index-1] && same {
			continue
		}
		if nums[i] != nums[index-1] {
			nums[index] = nums[i]
			same = false
			index++
		}
	}
	nums = nums[:index]
	return len(nums)
}
