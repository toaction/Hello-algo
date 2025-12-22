package skill

/*
287. 寻找重复数: https://leetcode.cn/problems/find-the-duplicate-number/

给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），
可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

示例 1：
	输入：nums = [1,3,4,2,2]
	输出：2
*/

// 数组中存在重复的整数 -> 链表中有环，链表头节点为 0
// 找到数组中重复的整数 -> 多个整数指向的同一个整数，即入环节点
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	// 找到相遇节点
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 找到入环节点
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
