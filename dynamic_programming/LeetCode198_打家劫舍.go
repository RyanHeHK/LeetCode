package dynamic_programming

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i-2]+nums[i] > nums[i-1] {
			res[i] = nums[i-2] + nums[i]
		} else {
			res[i] = nums[i-1]
		}
	}
	return res[len(nums)-1]
}
