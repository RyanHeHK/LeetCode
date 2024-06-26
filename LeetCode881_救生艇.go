package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left := 0
	right := len(people) - 1
	boats := 0
	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			left++
			right--
		}
		boats++
	}
	return boats
}
