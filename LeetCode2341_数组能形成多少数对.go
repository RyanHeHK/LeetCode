package main

func numberOfPairs(nums []int) []int {
	res := []int{}
	pair := 0
	vis := make([]bool, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if vis[i] == true {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if vis[j] == true {
				continue
			}
			if nums[i] == nums[j] {
				pair++
				vis[i], vis[j] = true, true
				break
			}
		}
	}
	res = append(res, pair)
	res = append(res, len(nums)-pair*2)
	return res
}
