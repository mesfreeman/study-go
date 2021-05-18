/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	// 基础映射
	baseMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {
		v1, ok1 := baseMap[string(s[i])]
		if !ok1 {
			return 0
		}

		if i+1 < len(s) {
			v2, ok2 := baseMap[string(s[i+1])]
			if !ok2 {
				return 0
			}

			// 前值大于等于后值
			if v1 >= v2 {
				result += v1
				continue
			}

			// 前值小于后值
			result += (v2 - v1)
			i += 1 // 跳过
		} else {
			result += v1
		}
	}
	return result
}

// @lc code=end

