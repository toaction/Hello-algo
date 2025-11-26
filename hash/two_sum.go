package hash

// 1. 两数之和
//
// https://leetcode.cn/problems/two-sum/
//
// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
// 示例 1：
//
//	输入：nums = [2,7,11,15], target = 9
//	输出：[0,1]
//	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
// 解题思路：
//	- 哈希表存储：key = 数字值，value = 对应索引
//	- 单次遍历：遍历数组时查找 complement = target - nums[i]
//	- 快速查找：如果 complement 在哈希表中存在，说明找到答案
//	- 逐步构建：如果 complement 不存在，将当前数字及其索引存入哈希表
//
// 时间复杂度：O(n) - 单次遍历数组
// 空间复杂度：O(n) - 最坏情况下需要存储所有元素
func TwoSum(nums []int, target int) []int {
	// key = num, value = idx
	ht := map[int]int{}
	for i, x := range nums {
		if p, ok := ht[target-x]; ok {
			return []int{p, i}
		}
		ht[x] = i
	}
	return nil
}
