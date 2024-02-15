package main

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	var majority int
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			majority = k
		}
	}
	return majority
}
