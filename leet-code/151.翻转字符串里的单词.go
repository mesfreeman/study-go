/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	// 以空格进行切分，然后倒序遍历
	res := ""
	strSlice := strings.Split(s, " ")
	for i := len(strSlice) - 1; i >= 0; i-- {
		tmp := strings.TrimSpace(strSlice[i])
		if tmp == "" {
			continue
		}
		res += " " + tmp
	}
	return strings.TrimSpace(res)
}

// @lc code=end

