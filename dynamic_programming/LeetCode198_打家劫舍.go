package dynamic_programming

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	d1 := nums[0]
	d2 := 0
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			if d2+nums[i] > d1 {
				d2 = d2 + nums[i]
			} else {
				d2 = d1
			}
		} else {
			if d1+nums[i] > d2 {
				tmp := d1 + nums[i]
				d1 = d2
				d2 = tmp
			} else {
				d1 = d2
			}
		}
	}
	return d2
}
