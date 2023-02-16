package main

func permute(nums []int) [][]int {
	var output [][]int
	var backtrack func(first int)
	backtrack = func(first int) {
		if first == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			output = append(output, tmp)
			return
		}
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return output
}
