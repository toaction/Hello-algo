package binarysearch

/*
35. 搜索插入位置：https://leetcode.cn/problems/search-insert-position/

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

示例 1:
	输入: nums = [1,3,5,6], target = 5
	输出: 2
*/

// 二分查找
// 找到第一个 大于等于target 的元素下标
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
