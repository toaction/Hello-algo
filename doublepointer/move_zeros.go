package doublepointer

// 283. 移动零
//
// https://leetcode.cn/problems/move-zeroes/
//
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例：
//
//	输入: nums = [0,1,0,3,12]
//	输出: [1,3,12,0,0]
//
// MoveZeroes 将所有0移动到数组末尾，同时保持非零元素的相对顺序
//
// 解题思路：
// 1. 使用快慢指针：left指针指向当前非零元素应该放置的位置，right指针遍历数组
// 2. 当right指针遇到非零元素时，将其与left指针位置交换，然后left指针后移
// 3. 当right指针遇到0时，只移动right指针，left指针保持不变
// 4. 这样所有非零元素都会按原有顺序被交换到数组前部，0自然被挤到后部
//
// 时间复杂度：O(n) - 只需遍历数组一次
// 空间复杂度：O(1) - 原地操作，不需要额外空间
func MoveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			// swap
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
