import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 *
 * https://leetcode-cn.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (51.27%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 11.6K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 * 给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold
 * Medal", "Silver Medal", "Bronze Medal"）。
 *
 * (注：分数越高的选手，排名越靠前。)
 *
 * 示例 1:
 *
 *
 * 输入: [5, 4, 3, 2, 1]
 * 输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * 解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal"
 * and "Bronze Medal").
 * 余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
 *
 * 提示:
 *
 *
 * N 是一个正整数并且不会超过 10000。
 * 所有运动员的成绩都不相同。
 *
 *
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	nums备份 := []int{}
	for _, i := range nums {
		nums备份 = append(nums备份, i)
	}
	res := []string{}
	名次 := make(map[int]int)
	sort.Ints(nums)
	key := len(nums)
	for _, value := range nums {
		名次[value] = key
		key--
	}
	标记 := 0
	for _, i := range nums备份 {
		switch 名次[i] {
		case 1:
			res = append(res, "Gold Medal")
			标记 = 1
		case 2:
			res = append(res, "Silver Medal")
			标记 = 1
		case 3:
			res = append(res, "Bronze Medal")
			标记 = 1
		}
		if 标记 == 0 {
			res = append(res, strconv.Itoa(名次[i]))
		}
		标记 = 0
	}
	return res
}

// @lc code=end
