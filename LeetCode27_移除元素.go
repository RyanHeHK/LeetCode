package main

func removeElement(nums []int, val int) int {
	l := len(nums)
	index := 0
	for i := range nums {
		if nums[i] != val {
			if index != i {
				nums[index] = nums[i]
			}
			index++
		} else {
			l--
		}
	}
	return l
}
