package normalArray

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
//
// 前缀乘积 * 后缀乘积
func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	cur := 1
	for i, v := range nums {
		ans[i] = cur
		cur = cur * v
	}
	cur = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * cur
		cur = cur * nums[i]
	}
	return ans
}
