/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (51.86%)
 * Likes:    282
 * Dislikes: 0
 * Total Accepted:    42.6K
 * Total Submissions: 82.1K
 * Testcase Example:  '1'
 *
 * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
 * 
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 * 
 * 
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 * 
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
 * 
 * 注意：整数顺序将表示为一个字符串。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: 1
 * 输出: "1"
 * 
 * 
 * 示例 2:
 * 
 * 输入: 4
 * 输出: "1211"
 * 
 * 
 */
 func countAndSay(n int) string {
	str := "1"
	for i:=1;i<n;i++{
		res := ""
		number:=str[0]
		count:=0
		for j:=0;j<len(str);j++{
			if str[j] == number {
				count++
			 } else {
				res = res + strconv.Itoa(count) + string(number) 
				number = str[j]
				count = 1
			 }
		}
		res = res + strconv.Itoa(count) + string(number) 
		str = res
	}
	return str
}

