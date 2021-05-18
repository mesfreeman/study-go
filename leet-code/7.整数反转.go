/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var result int
	for x != 0 {
		y := x % 10
		result = result*10 + y
		x = x / 10
	}

	maxInt := 1<<31 - 1
	minInt := -1 << 31
	if maxInt < result || minInt > result {
		return 0
	}

	return result
}

// @lc code=end

