/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var result int
	for x != 0 {
		tmp := x % 10
		result = result*10 + tmp
		x = x / 10
	}

	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	if result > MaxInt32 || result < MinInt32 {
		return 0
	}

	return result
}

// @lc code=end

