package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || tomatoSlices > cheeseSlices*4 || tomatoSlices < cheeseSlices*2 {
		return []int{}
	}
	return []int{(tomatoSlices - 2*cheeseSlices) / 2, (4*cheeseSlices - tomatoSlices) / 2}
}
