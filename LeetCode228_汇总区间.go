package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	left := 0
	right := 1
	if len(nums) == 0 {
		return res
	}
	for right < len(nums) {
		if nums[right]-nums[right-1] == 1 {
			right++
		} else {
			if nums[right-1] == nums[left] {
				res = append(res, strconv.Itoa(nums[left]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[right-1]))
			}
			left = right
			right++
		}
	}
	if nums[right-1] == nums[left] {
		res = append(res, strconv.Itoa(nums[left]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[right-1]))
	}
	return res
}
