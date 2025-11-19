package dp

// 279. 完全平方数
//
// https://leetcode.cn/problems/perfect-squares/
//
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
// 示例 1：
//
//	输入：n = 12
//	输出：3
//	解释：12 = 4 + 4 + 4
func NumSquares(n int) int {
	// dp[i]：和为 i 的完全平方数的最少数量
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// i 个 1 相加
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
