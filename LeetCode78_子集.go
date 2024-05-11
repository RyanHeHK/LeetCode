package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	deep := 0
	var dfs func(deep int, index int, subArray []int, res [][]int) [][]int
	dfs = func(deep int, index int, subArray []int, res [][]int) [][]int {
		if deep == len(subArray) {
			tmp := make([]int, len(subArray))
			copy(tmp, subArray)
			res = append(res, tmp)
			return res
		}
		for i := index; i < len(nums); i++ {
			subArray = append(subArray, nums[i])
			res = dfs(deep, i+1, subArray, res)
			subArray = subArray[:len(subArray)-1]
		}
		return res
	}

	for deep <= len(nums) {
		subArray := []int{}
		res = append(res, dfs(deep, 0, subArray, [][]int{})...)
		deep++
	}
	return res
}
