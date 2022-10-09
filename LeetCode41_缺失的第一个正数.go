package main

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		m[nums[i]] = true
	}

	for i := 1; i <= max; i++ {
		if !m[i] {
			return i
		}
	}
	return max + 1
}
