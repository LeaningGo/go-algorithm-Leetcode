/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (46.44%)
 * Likes:    5960
 * Dislikes: 0
 * Total Accepted:    500.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 
 * 示例:
 * 
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 * 
 * 
 */
 func twoSum(nums []int, target int) []int {
	var a []int
	b := 0
	for key, i := range nums {
		for key2, j := range nums {
			if (target == (i + j)) && (key != key2) {
				a = append(a, key)
				a = append(a, key2)
				b = 1
				break
			}
		}
		if b == 1 {
			break
		}
	}
	return a
}

