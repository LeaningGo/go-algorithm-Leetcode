/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.42%)
 * Likes:    651
 * Dislikes: 0
 * Total Accepted:    108.3K
 * Total Submissions: 314.8K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 示例 1:
 * 
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 * 
 * 
 * 说明:
 * 
 * 所有输入只包含小写字母 a-z 。
 * 
 */
 func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lengthmin := len(strs[0])
	var min string
	for _, i := range strs {
		if len(i) <= lengthmin {
			lengthmin = len(i)
			min = i
		}
	}
	var minstring []string

	for i := len(min); i > 0; i-- {
		minstring = append(minstring, min[0:i])
	}
	for _, i := range minstring {
		flag := 0
		for _, j := range strs {
			if i != j[0:len(i)] {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return i
		}
	}
	return ""
}

