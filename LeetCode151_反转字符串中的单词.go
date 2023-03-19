package main

func reverseWords(s string) string {
	m := []string{}
	wLeft := false
	left := 0
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && !wLeft {
			wLeft = true
			left = i
		}
		if s[i] == ' ' && wLeft {
			m = append(m, s[left:i])
			wLeft = false
		}
	}
	if wLeft {
		m = append(m, s[left:])
	}
	for i := len(m) - 1; i >= 0; i-- {
		if i != len(m)-1 {
			res += " "
		}
		res += m[i]
	}
	return res
}
