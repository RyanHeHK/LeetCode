package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	if purchaseAmount == 0 {
		return 100
	}
	payAmount := 0
	y := purchaseAmount % 10
	if y < 5 {
		payAmount = purchaseAmount / 10 * 10
	} else {
		payAmount = (purchaseAmount/10 + 1) * 10
	}

	return 100 - payAmount
}
