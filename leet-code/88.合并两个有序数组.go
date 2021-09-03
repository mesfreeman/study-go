/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 没有元素要合并，直接返回
	if n == 0 {
		return
	}

	// 先填充 num2 元素到 num1中
	for _, val := range nums2 {
		nums1[m] = val
		m++
	}

	// 冒泡排序
	len := len(nums1)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j+1], nums1[j] = nums1[j], nums1[j+1]
			}
		}
	}
}

// @lc code=end

