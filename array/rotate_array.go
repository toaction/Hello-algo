package array

/*
189. 轮转数组

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
	输入: nums = [1,2,3,4,5,6,7], k = 3
	输出: [5,6,7,1,2,3,4]
	解释:
		向右轮转 1 步: [7,1,2,3,4,5,6]
		向右轮转 2 步: [6,7,1,2,3,4,5]
		向右轮转 3 步: [5,6,7,1,2,3,4]
*/

// 先整体反转，然后左部分和右部分分别反转
// -->--->
// <---<--
// --->-->
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
