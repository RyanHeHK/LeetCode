package main

func maxChunksToSorted(arr []int) int {
	tmp := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > tmp {
			tmp = arr[i]
		}
		if tmp == i {
			count++
		}
	}
	return count
}
