package greed

// 45. 跳跃游戏 II
//
// https://leetcode.cn/problems/jump-game-ii/
//
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。
//
// 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：
//	- 0 <= j <= nums[i] 且
//	- i + j < n
// 返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。
//
// 示例 1：
//  输入：nums = [2,3,1,1,4]
//  输出：2
//  解释：
//      跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func Jump(nums []int) int {
	ans := 0
	// 本轮区间终点
	end := 0
	// 以本轮区间中的点为起点，所能到达的最远距离
	access := 0
	for i, v := range nums {
		// 更新区间
		if i > end {
			ans++
			end = access
		}
		access = max(access, i+v)
	}
	return ans
}
