package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	size := len(nums)
	m := make([]int, size)
	m[0] = nums[0]
	for i := 1; i < size; i++ {
		m[i] = m[i-1] + nums[i]
	}
	res := []int{}
	for _, v := range queries {
		res = append(res, binarySearch(m, v, 0, size-1)+1)
	}
	return res
}

func binarySearch(m []int, v, left, right int) int {
	if left == right {
		if v < m[left] {
			return left - 1
		}
		return left
	}
	mid := (left + right) / 2
	if m[mid] <= v {
		left = mid + 1
	}
	if v < m[mid] {
		right = mid
	}
	return binarySearch(m, v, left, right)
}
