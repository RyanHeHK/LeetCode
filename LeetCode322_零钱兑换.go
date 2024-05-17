package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	count := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		count[i] = -1
	}
	for i := 0; i < amount+1; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				if count[i-c] == -1 {
					continue
				} else {
					if count[i] == -1 {
						count[i] = count[i-c] + 1
					} else {
						count[i] = min(count[i], count[i-c]+1)
					}
				}
			}
		}
	}
	return count[amount]
}
