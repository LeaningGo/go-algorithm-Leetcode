/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (32.93%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    13.5K
 * Total Submissions: 40.7K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [3, 2, 1]
 * 
 * 输出: 1
 * 
 * 解释: 第三大的数是 1.
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: [1, 2]
 * 
 * 输出: 2
 * 
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: [2, 2, 3, 1]
 * 
 * 输出: 1
 * 
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 * 
 * 
 */

// @lc code=start
func thirdMax(nums []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	一, 二, 三 := INT_MIN, INT_MIN, INT_MIN
	for i := 0; i < len(nums); i++ {
		if nums[i] > 一 {
			三 = 二
			二 = 一
			一 = nums[i]
		} else if nums[i] > 二 && nums[i] != 一 {
			三 = 二
			二 = nums[i]
		} else if nums[i] > 三 && nums[i] != 二 && nums[i] != 一 {
			三 = nums[i]
		}
	}
	if 二 == INT_MIN || 三 == INT_MIN {
		return 一
	}
	fmt.Println(一, 二, 三)
	return 三
}
// @lc code=end

