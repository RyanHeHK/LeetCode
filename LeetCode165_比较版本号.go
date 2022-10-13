package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	res := 0
	index := 0
	l1 := len(v1)
	l2 := len(v2)
	if l1 < l2 {
		index = l1
		res = compareSubVersion(v1, v2)
		if res != 0 {
			return res
		}
		for i := index; i < l2; i++ {
			if subV, _ := strconv.Atoi(v2[i]); subV > 0 {
				return -1
			}
		}
	} else {
		index = l2
		res = compareSubVersion(v2, v1)
		if res != 0 {
			return -res
		}
		for i := index; i < l1; i++ {
			if subV, _ := strconv.Atoi(v1[i]); subV > 0 {
				return 1
			}
		}
	}
	return res
}

func compareSubVersion(v1, v2 []string) int {
	for i, _ := range v1 {
		sub1, _ := strconv.Atoi(v1[i])
		sub2, _ := strconv.Atoi(v2[i])
		if sub1 > sub2 {
			return 1
		}
		if sub1 < sub2 {
			return -1
		}
	}
	return 0
}
