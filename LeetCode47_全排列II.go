package main

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

func permuteUnique1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	mark := make([]bool, len(nums))
	var combination []int
	var dfs func(int)
	dfs = func(l int) {
		if l == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, combination)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if mark[i] || i > 0 && nums[i] == nums[i-1] && !mark[i-1] {
				continue
			}
			combination = append(combination, nums[i])
			mark[i] = true
			dfs(l + 1)
			mark[i] = false
			combination = combination[:len(combination)-1]
		}
	}
	dfs(0)
	return res
}
