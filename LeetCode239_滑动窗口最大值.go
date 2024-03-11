package main

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	res := []int{}
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := k; i < len(nums); i++ {
		res = append(res, nums[q[0]])
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-k >= q[0] {
			q = q[1:]
		}
	}
	res = append(res, nums[q[0]])
	return res
}
