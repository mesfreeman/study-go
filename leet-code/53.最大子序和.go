/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 动态规则
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果和前值加和之后变大，则当前值改为加和之后的值
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

// @lc code=end

