package doublepointer

/* 
283. 移动零：https://leetcode.cn/problems/move-zeroes/

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例：
	输入: nums = [0,1,0,3,12]
	输出: [1,3,12,0,0] 
*/


// 双指针，left 指向非零元素要移动的位置
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
