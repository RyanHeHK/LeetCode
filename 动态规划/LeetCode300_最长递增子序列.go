package 动态规划

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxAns := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if maxAns < dp[i] {
			maxAns = dp[i]
		}
	}
	return maxAns
}
