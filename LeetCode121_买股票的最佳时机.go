package main

import "LeetCode/utils"

// 贪心
func maxProfit121(prices []int) int {
	buySell := make([]int, 2)
	res := 0
	buySell[0] = prices[0]
	for i := range prices {
		if prices[i] >= buySell[0] {
			buySell[1] = prices[i]
			res = utils.Max(res, buySell[1]-buySell[0])
		} else {
			buySell[0] = prices[i]
			buySell[1] = prices[i]
		}
	}
	return res
}
