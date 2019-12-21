/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (50.54%)
 * Likes:    244
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 77.3K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
func addBinary(a string, b string) string {
	c := ""
	l, s := "", ""
	if len(a) > len(b) {
		l = a
		s = b
	} else {
		l = b
		s = a
	}
	cha := len(l) - len(s)
	bu := ""
	for i := 0; i < cha; i++ {
		bu = bu + "0"
	}
	s = bu + s
	var wei uint8
	for i := len(l) - 1; i >= 0; i-- {
		tmp := (l[i] - '0') + (s[i] - '0') + wei
		if tmp >= 2 {
			wei = 1
			c = string(tmp-2+'0') + c
		} else {
			c = string(tmp+'0') + c
			wei = 0
		}
	}
	if wei == byte(1) {
		c = "1"+ c
	}
	return c
}

