package main

func maxProfit1(prices []int) int {
	remain1, remain2 := -prices[0], -prices[0]
	profit1, profit2 := 0, 0
	for i := 1; i < len(prices); i++ {
		if remain1 < -prices[i] {
			remain1 = -prices[i]
		}
		if profit1 < prices[i]+remain1 {
			profit1 = prices[i] + remain1
		}
		if remain2 < profit1-prices[i] {
			remain2 = profit1 - prices[i]
		}
		if profit2 < prices[i]+remain2 {
			profit2 = prices[i] + remain2
		}
	}
	return profit2
}
