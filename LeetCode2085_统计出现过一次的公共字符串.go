package main

func countWords(words1 []string, words2 []string) int {
	w1map := make(map[string]int)
	w2map := make(map[string]int)
	count := 0
	for _, w1 := range words1 {
		w1map[w1]++
	}
	for _, w2 := range words2 {
		w2map[w2]++
	}
	for k, v := range w1map {
		if v == 1 && w2map[k] == 1 {
			count++
		}
	}
	return count
}
