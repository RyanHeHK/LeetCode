package main

// https://leetcode.cn/problems/longest-palindromic-substring/
// 中心法
func longestPalindrome(s string) string {
	maxL := 0
	begin := 0
	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1
		L := 1
		for true {
			if left >= 0 && s[left] == s[i] {
				L++
				left--
			} else {
				break
			}
		}
		for true {
			if right < len(s) && s[right] == s[i] {
				L++
				right++
			} else {
				break
			}
		}
		for true {
			if left >= 0 && right < len(s) && s[left] == s[right] {
				L += 2
				left--
				right++
			} else {
				break
			}
		}
		if L > maxL {
			maxL = L
			begin = left + 1
		}
	}
	return s[begin : begin+maxL]
}

// dynamic_programming
func longestPalindrome1(s string) string {
	l := len(s)
	dp := [][]bool{}
	for i := 0; i < l; i++ {
		subDp := []bool{}
		for j := 0; j < l; j++ {
			subDp = append(subDp, true)
		}
		dp = append(dp, subDp)
	}
	maxL := 1
	begin := 0
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && maxL < j-i+1 {
				maxL = j - i + 1
				begin = i
			}
		}

	}
	return s[begin : begin+maxL]
}
