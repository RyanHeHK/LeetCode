package main

func maxProfit122(prices []int) int {
	res := 0
	buyPrice := prices[0]
	sellPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < sellPrice {
			res += sellPrice - buyPrice
			buyPrice = prices[i]
			sellPrice = prices[i]
			continue
		}
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		}
		if prices[i] > sellPrice {
			sellPrice = prices[i]
		}
	}
	res += sellPrice - buyPrice
	return res
}
