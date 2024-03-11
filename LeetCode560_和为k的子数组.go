package main

func subarraySum(nums []int, k int) int {
	subArraySum := map[int]int{}
	tmpN := 0
	res := 0
	subArraySum[0] = 1
	for _, n := range nums {
		tmpN += n
		subArraySum[tmpN]++
		if subArraySum[n-k] != 0 {
			res += subArraySum[n-k]
		}
	}
	return res
}
