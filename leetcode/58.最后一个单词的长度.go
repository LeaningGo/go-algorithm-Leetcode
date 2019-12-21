/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (30.72%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 126.3K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0
	}
	str := ""
	back := len(s) - 1
	for i := back; i >= 0; i-- {
		if string(s[i]) != " " {
			break
		} else {
			back--
		}
	}
	for i := back; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		}
		str = str + string(s[i])
	}
	return len(str)
}
