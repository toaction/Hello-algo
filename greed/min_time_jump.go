package greed

/*
45. 跳跃游戏 II: https://leetcode.cn/problems/jump-game-ii/

给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。
每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：
  - 0 <= j <= nums[i] 且
  - i + j < n
返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。

示例 1：
	输入：nums = [2,3,1,1,4]
	输出：2
	解释：
	    跳到最后一个位置的最小跳跃数是 2。
	    从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*/

// 每次跳跃都从一个区间跳到另一个区间，下一个区间位置可从上一个区间位置跳跃得到
// 上一个区间为 [i, j]
// 下一个区间的范围为 [j+1, rightmost], rightmost 为从 [i, j]区间位置所能跳跃到的最远距离
func Jump(nums []int) int {
	ans := 0
	right, rightmost := 0, 0
	for i, v := range nums {
		if i > right {
			// 从一个区间跳到下一个区间，跳跃数 + 1
			ans++
			right = rightmost
		}
		rightmost = max(rightmost, i+v)
	}
	return ans
}
