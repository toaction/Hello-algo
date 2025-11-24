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
