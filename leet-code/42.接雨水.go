/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	ans := 0                 // 总雨水数
	maxIdx := 0              // 最高元素索引
	maxVal := height[maxIdx] // 最高元素值
	len := len(height)

	if len == 0 {
		return ans
	}

	// 找出最高元素
	for i := 0; i < len; i++ {
		if maxVal <= height[i] {
			maxIdx = i
			maxVal = height[i]
		}
	}

	// 从左到右遍历，如果左边元素大于右边元素则表示可以接到水
	for left := 0; left < maxIdx; left++ {
		for j := left + 1; j <= maxIdx; j++ {
			if height[left] > height[j] {
				ans += height[left] - height[j]
			} else {
				left = j
			}
		}
	}

	// 从右向左遍历，如果右边元素大于左边元素则表示可以接到水
	for right := len - 1; right > maxIdx; right-- {
		for k := right - 1; k >= maxIdx; k-- {
			if height[right] > height[k] {
				ans += height[right] - height[k]
			} else {
				right = k
			}
		}
	}

	return ans
}

// @lc code=end

