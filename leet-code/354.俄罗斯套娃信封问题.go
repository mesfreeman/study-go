/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	len := len(envelopes)
	if len == 1 {
		return 1
	}

	// 先冒泡排序
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			// 先以第一个元素大小为准
			if envelopes[j][0] >= envelopes[j+1][0] {
				envelopes[j+1], envelopes[j] = envelopes[j], envelopes[j+1]
			}

			// 如果第一个元素相等，则对比第二个元素
			if envelopes[j][0] == envelopes[j+1][0] && envelopes[j][1] >= envelopes[j+1][1] {
				envelopes[j+1], envelopes[j] = envelopes[j], envelopes[j+1]
			}
		}
	}

	// 按递增子序列进行处理
	dp := make([]int, len)
	dp[0] = 1
	max := 1
	for i := 1; i < len; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] && dp[j]+1 > dp[i] {
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

