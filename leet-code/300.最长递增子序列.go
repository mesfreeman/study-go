/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	dp[0] = 1
	max := 1

	for i := 1; i < len; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 如果前值比后值大，刚前值在后值最大子序列的基础上加1
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// @lc code=end

