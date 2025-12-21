package binarysearch

/*
34. 在排序数组中查找元素的第一个和最后一个位置: https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

示例 1：
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]
*/

// 找到第一个小于target的元素和第一个大于target的元素
func SearchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return ans
	}
	ans[0] = left
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	ans[1] = right
	return ans
}
