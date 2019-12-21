/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (42.15%)
 * Likes:    76
 * Dislikes: 0
 * Total Accepted:    18.7K
 * Total Submissions: 43.9K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 0; i <= num/2; i++ {
		if num == i*i {
			return true
		}
	}
	return false
}

// @lc code=end
