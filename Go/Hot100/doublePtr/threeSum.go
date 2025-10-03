package doublePtr

import (
	"sort"
)

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 思路：
//
// 1. 先排序
//
// 2. 使用三指针，i,left,right 指向三个数，
// i 从0开始，left 从 i+1 开始，right 从 len(nums)-1 开始,
// 若和小于 0，则 left 右移，若和大于 0，则 right 左移，若和等于 0，则将当前的三个数加入到结果集中
//
// 3. 去重:
// 若 i>0 && nums[i]==nums[i-1]，则跳过，因为已经遍历过,
// 若 left>i && nums[left]==nums[left-1]，则跳过，因为已经遍历过,
// 不用对 right 去重，因为 left 去重后，结果一定会改变
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	idx := 0
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
				idx++
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
