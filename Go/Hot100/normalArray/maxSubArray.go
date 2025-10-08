package normalArray

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组是数组中的一个连续部分。
//
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
// 前缀和
func MaxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	minPreSum := 0
	for _, v := range nums {
		sum = sum + v
		ans = max(ans, sum-minPreSum)
		minPreSum = min(minPreSum, sum)
	}
	return ans
}
