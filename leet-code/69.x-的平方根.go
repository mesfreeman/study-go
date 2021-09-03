/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	res := 0
	for i := 1; i <= x/2; i++ {
		if i*i <= x {
			res = i
		} else {
			break
		}
	}
	return res
}

// @lc code=end

