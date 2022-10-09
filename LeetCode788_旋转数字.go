package main

func rotatedDigits(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		num := i + 1
		valid := false
		for num > 0 {
			yu := num % 10
			if yu == 2 || yu == 5 || yu == 6 || yu == 9 {
				valid = true
			}
			if yu == 3 || yu == 4 || yu == 7 {
				valid = false
				break
			}
			num = num / 10
		}
		if valid {
			count++
		}
	}
	return count
}
