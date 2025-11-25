package doublepointer

import "sort"

// 15. 三数之和
//
// https://leetcode.cn/problems/3sum/
//
// 给你一个整数数组 nums ，
// 判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。
// 请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
//
//	输入：nums = [-1,0,1,2,-1,-4]
//	输出：[[-1,-1,2],[-1,0,1]]
//	解释：
//	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//	不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//
// 解题思路：
// 1. 先排序：使数组有序，便于去重和使用双指针
// 2. 固定第一个元素nums[i]，然后用双指针在剩余数组中寻找另外两个数
// 3. 双指针查找：left=i+1, right=n-1
//   - 如果和小于0：移动left指针增大和
//   - 如果和大于0：移动right指针减小和
//   - 如果和等于0：找到一组解，同时需要去重处理
//
// 4. 去重策略：
//   - 外层去重：跳过nums[i] == nums[i-1]的情况
//   - 内层去重：找到解后，跳过相同的nums[left]值
//
// 时间复杂度：O(n²) - 排序O(nlogn) + 双重循环O(n²)
// 空间复杂度：O(1) - 不考虑返回结果的话
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去除多余元素
				tmp := nums[left]
				for left < right && nums[left] == tmp {
					left++
				}
			}
		}
	}
	return ans
}
