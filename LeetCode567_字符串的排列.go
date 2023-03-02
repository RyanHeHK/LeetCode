package main

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	for s1Len > s2Len {
		return false
	}
	count := [26]int{}
	for _, c := range s1 {
		count[c-'a']++
	}
	left, right := 0, 0
	for right < s2Len {
		count[s2[right]-'a']--
		for count[s2[right]-'a'] < 0 {
			count[s2[left]-'a']++
			left++
		}
		if right-left+1 == s1Len {
			return true
		}
		right++
	}
	return false
}
