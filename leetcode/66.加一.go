/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.30%)
 * Likes:    330
 * Dislikes: 0
 * Total Accepted:    72K
 * Total Submissions: 177.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */
func plusOne(digits []int) []int {
	for long := len(digits) - 1; long >= 0; long-- {
		if digits[long]+1 < 10 {
			digits[long]++
			break
		} else {
			digits[long] = 0
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
		return digits
	} else {
		return digits
	}
}
