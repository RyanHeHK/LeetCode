package main

//func minWindow(s string, t string) string {
//	sLen, tLen := len(s), len(t)
//	if tLen > sLen {
//		return ""
//	}
//	left, right := 0, 0
//	match := 0
//	minSubStr := ""
//	count := make(map[byte]int)
//	for i := 0; i < tLen; i++ {
//		count[t[i]]++
//		match++
//	}
//
//	for right < sLen {
//		if count[s[right]] > 0 {
//			match--
//		}
//		count[s[right]]--
//
//		for count[s[left]] < 0 && left < right {
//			count[s[left]]++
//			left++
//		}
//		if match == 0 {
//			if len(minSubStr) == 0 {
//				minSubStr = s[left : right+1]
//			} else {
//				if right-left+1 < len(minSubStr) {
//					minSubStr = s[left : right+1]
//				}
//			}
//			for left < right {
//				if match == 1 {
//					break
//				}
//				if count[s[left]] == 0 {
//					match++
//				}
//				count[s[left]]++
//				left++
//			}
//		}
//		right++
//	}
//	return minSubStr
//}

func minWindow(s string, t string) string {
	m := map[int32]int{}
	diff := 0
	for _, c := range t {
		m[c-'A']++
		diff++
	}
	l, r := 0, 0
	res := ""
	for _, c := range s {
		r++
		if _, ok := m[c-'A']; ok {
			m[c-'A']--
			if m[c-'A'] >= 0 {
				diff--
			}
		}
		for diff == 0 && l < r {
			res = minString(res, s[l:r])
			if _, ok := m[int32(s[l]-'A')]; ok {
				m[int32(s[l]-'A')]++
				if m[int32(s[l]-'A')] > 0 {
					diff++
				}
			}
			l++
		}
	}
	return res
}

func minString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	if len(s1) < len(s2) {
		return s1
	} else {
		return s2
	}
}
