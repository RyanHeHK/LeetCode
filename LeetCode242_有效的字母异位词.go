package main

func isAnagram(s string, t string) bool {
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a'] += 1
	}
	for _, c := range t {
		m[c-'a'] -= 1
	}
	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}
	return true
}
