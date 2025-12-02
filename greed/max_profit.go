package greed

/*
121. 买卖股票的最佳时机: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：
	输入：prices = [7,1,5,3,6,4]
	输出：5
	解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
*/

// 贪心
// 使用一个数记录今天及之前股票的最低价
// 假设最低价那天买入，今天卖出
func MaxProfit(prices []int) int {
	ans := 0
	low := prices[0]
	for _, v := range prices {
		low = min(low, v)
		ans = max(ans, v-low)
	}
	return ans
}
