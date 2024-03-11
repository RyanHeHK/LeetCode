package main

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	emptyBottles := numBottles
	for emptyBottles/numExchange > 0 {
		res = res + emptyBottles/numExchange
		emptyBottles = emptyBottles/numExchange + emptyBottles%numExchange
	}
	return res
}
