/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	best := 0
	for i := len(prices)-1; i > 0; i-- {
		for j := i-1; j >= 0; j-- {
			if prices[i]-prices[j] > best {
				best = prices[i] - prices[j]
			}
		}
	}
	return best
}

// @lc code=end
