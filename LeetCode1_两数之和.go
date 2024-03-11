package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	res := []int{}
	for i, n := range nums {
		if _, ok := m[target-n]; ok {
			res = append(res, m[target-n], i)
		}
		m[n] = i
	}
	return res
}
