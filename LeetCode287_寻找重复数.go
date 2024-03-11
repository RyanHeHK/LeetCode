package main

func findDuplicate(nums []int) int {
	m := map[int]int{}
	for i := range nums {
		if m[nums[i]] == 1 {
			return nums[i]
		}
		m[nums[i]]++
	}
	return -1
}
