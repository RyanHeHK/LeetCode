package main

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if tLen > sLen {
		return ""
	}
	left, right := 0, 0
	match := 0
	minSubStr := ""
	count := make(map[byte]int)
	for i := 0; i < tLen; i++ {
		count[t[i]]++
		match++
	}

	for right < sLen {
		if count[s[right]] > 0 {
			match--
		}
		count[s[right]]--

		for count[s[left]] < 0 && left < right {
			count[s[left]]++
			left++
		}
		if match == 0 {
			if len(minSubStr) == 0 {
				minSubStr = s[left : right+1]
			} else {
				if right-left+1 < len(minSubStr) {
					minSubStr = s[left : right+1]
				}
			}
			for left < right {
				if match == 1 {
					break
				}
				if count[s[left]] == 0 {
					match++
				}
				count[s[left]]++
				left++
			}
		}
		right++
	}
	return minSubStr
}
