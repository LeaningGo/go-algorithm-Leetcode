/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (30.50%)
 * Likes:    235
 * Dislikes: 0
 * Total Accepted:    35K
 * Total Submissions: 112.2K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 * 
 * 示例:
 * 
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 * 
 * 
 */

// @lc code=start
func countPrimes(n int) int {
	if n == 0 || n == 1{
		return 0
	}
	isPrime:=make([]int,n)
	for i:=2;i*i<n;i++{
		if isPrime[i]==0{
			for j:=i*i;j<n;j=j+i{
				isPrime[j]=1
			}
		}
	}
	count :=0
	for _,i := range isPrime {
		if i == 0{
			count++
		}
	}
	return count-2
}
// @lc code=end

