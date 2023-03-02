package main

func findAnagrams(s string, p string) (res []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}
	pos := make([]int, 26)
	for i := 0; i < n; i++ {
		pos[p[i]-'a']++
	}
	for l, r := 0, 0; r < m; r++ {
		pos[s[r]-'a']--

		for pos[s[r]-'a'] < 0 {
			pos[s[l]-'a']++
			l++
		}
		if r-l+1 == n {
			res = append(res, l)
		}
	}
	return
}
