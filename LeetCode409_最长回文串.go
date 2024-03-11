package main

func longestPalindrome2(s string) int {
	cMap := map[byte]int{}
	for i := range s {
		cMap[s[i]-'a']++
	}
	res := 0
	oddNum := false
	for i := range cMap {
		if cMap[i]%2 == 0 {
			res += cMap[i]
		}
		if cMap[i]%2 == 1 {
			res += cMap[i] - 1
			oddNum = true
		}
	}
	if oddNum {
		res++
	}
	return res
}
