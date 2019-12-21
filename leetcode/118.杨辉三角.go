/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (64.27%)
 * Likes:    192
 * Dislikes: 0
 * Total Accepted:    40.9K
 * Total Submissions: 63.7K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	triangle := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		rows := []int{1}
		for j := 1; j <= i-1; j++ {
			num := triangle[i-1][j-1] + triangle[i-1][j]
			rows = append(rows, num)
		}
		rows = append(rows, 1)
		triangle = append(triangle, rows)
	}
	return triangle
}

// @lc code=end

