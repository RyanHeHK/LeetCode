package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		return false
	})
	commonBallon := [][]int{}
	commonBallon = append(commonBallon, []int{points[0][0], points[0][1]})
	index := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] <= commonBallon[index][1] {
			commonBallon[index][0] = points[i][0]
			if commonBallon[index][1] > points[i][1] {
				commonBallon[index][1] = points[i][1]
			}
		} else {
			commonBallon = append(commonBallon, []int{points[i][0], points[i][1]})
			index++
		}
	}
	return len(commonBallon)
}
