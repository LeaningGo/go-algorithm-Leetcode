/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (40.96%)
 * Likes:    150
 * Dislikes: 0
 * Total Accepted:    46.7K
 * Total Submissions: 113.5K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 * 
 * 案例:
 * 
 * 
 * s = "leetcode"
 * 返回 0.
 * 
 * s = "loveleetcode",
 * 返回 2.
 * 
 * 
 * 
 * 
 * 注意事项：您可以假定该字符串只包含小写字母。
 * 
 */

// @lc code=start
func firstUniqChar(s string) int {
	a := [26]int{0}
	res := []int{}
	for _, i := range s {
		a[i-97]++
	}
	for m, i := range a {
		if i == 1 {
			n := string(m + 97)
			res = append(res, strings.Index(s, n))
		}
	}
	if len(res) == 0 {
		return -1
	}
	sort.Ints(res)
	return res[0]
}


// @lc code=end

