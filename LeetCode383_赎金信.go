package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	total := 0
	for _, c := range ransomNote {
		m[c-'a'] += 1
		total++
	}
	for _, c := range magazine {
		if total == 0 {
			return true
		}
		if m[c-'a'] == 0 {
			continue
		} else {
			m[c-'a'] -= 1
			total--
		}
	}
	if total == 0 {
		return true
	} else {
		return false
	}
}
