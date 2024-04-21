package 动态规划

import "LeetCode/utils"

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	count := make([]int, amount+1)
	for i, _ := range count {
		if i != 0 {
			count[i] = -1
		}
	}
	for i := 0; i < len(count); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if count[i-coins[j]] != -1 {
					if count[i] == -1 {
						count[i] = count[i-coins[j]] + 1
					} else {
						count[i] = utils.Min(count[i], count[i-coins[j]]+1)
					}
				}
			}
		}
	}
	return count[amount]
}
