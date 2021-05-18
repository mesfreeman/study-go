/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	len := len(str)

	for i := 0; i < len/2; i++ {
		if str[i] != str[len-i-1] {
			return false
		}
	}

	return true
}

// @lc code=end

