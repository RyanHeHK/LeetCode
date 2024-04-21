package main

import "LeetCode/utils"

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if numSet[num-1] {
			continue
		}
		count := 0
		for numSet[num] {
			num++
			count++
		}
		longestStreak = utils.Max(longestStreak, count)
	}
	return longestStreak
}
