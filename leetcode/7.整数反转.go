/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.82%)
 * Likes:    1231
 * Dislikes: 0
 * Total Accepted:    164.6K
 * Total Submissions: 501.7K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 * 
 * 示例 1:
 * 
 * 输入: 123
 * 输出: 321
 * 
 * 
 * 示例 2:
 * 
 * 输入: -123
 * 输出: -321
 * 
 * 
 * 示例 3:
 * 
 * 输入: 120
 * 输出: 21
 * 
 * 
 * 注意:
 * 
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 * 
 */
 func reverse(x int) int {
	var x4 int
	if x > 0 {
		x1 := strconv.Itoa(x)
		x2 := []rune(x1)
		for i, j := 0, len(x2)-1; i < j; i, j = i+1, j-1 {
			x2[i], x2[j] = x2[j], x2[i]
		}
		x3 := string(x2)
		x4, _ = strconv.Atoi(x3)
	} else {
		x := 0 - x
		x1 := strconv.Itoa(x)
		x2 := []rune(x1)
		for i, j := 0, len(x2)-1; i < j; i, j = i+1, j-1 {
			x2[i], x2[j] = x2[j], x2[i]
		}
		x3 := string(x2)
		x4, _ = strconv.Atoi(x3)
		x4 = 0 - x4
	}
	if x4 >= int(math.Pow(2, 31)-1) || x4 <= int(math.Pow(-2, 31)) {
		return 0
	} else {
		return x4
	}
}

