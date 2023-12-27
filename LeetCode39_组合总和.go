package main

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	var dfs func(candidates, combine []int, target, index, total int)
	dfs = func(candidates, combine []int, target, index, total int) {
		if target == total {
			if len(combine) == 0 {
				return
			}
			ans = append(ans, append([]int{}, combine...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if target-total-candidates[i] >= 0 {
				combine = append(combine, candidates[i])
				dfs(candidates, combine, target, i, total+candidates[i])
				combine = combine[:len(combine)-1]
			}
		}
	}
	dfs(candidates, []int{}, target, 0, 0)
	return ans
}
