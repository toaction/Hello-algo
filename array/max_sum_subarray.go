package array

/*
53. 最大子数组和: https://leetcode.cn/problems/maximum-subarray/

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：
	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

// 前缀和
// 使用 当前前缀和 减去 之前前缀和的最低点，得到的是以当前元素结尾的连续子数组的最大和
func MaxSubArrayBasePreSum(nums []int) int {
	ans := nums[0]
	sum, minsum := 0, 0
	for _, v := range nums {
		sum = sum + v
		ans = max(ans, sum-minsum)
		minsum = min(minsum, sum)
	}
	return ans
}

// 动态规划
// 令 f(i) 为以第 i 个数结尾的「连续子数组的最大和」，
// f(i) = max(f(i-1) + nums[i], nums[i])，加入前面子数组或单独成为子数组
func MaxSubArrayBaseDP(nums []int) int {
	ans := nums[0]
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		ans = max(ans, pre)
	}
	return ans
}
