package hash

// 128. 最长连续序列
//
// https://leetcode-cn.com/problems/longest-consecutive-sequence/
//
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 示例 1：
//
//	输入：nums = [100,4,200,1,3,2]
//	输出：4
//	解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 解题思路：
//	- 哈希去重：将数组转换为set，便于O(1)时间查找和去重
//	- 起始点判断：对每个数字，如果不存在x-1，说明x是连续序列的起始点
//	- 序列扩张：从起始点开始，不断查找x+1, x+2, x+3...直到中断
//	- 长度统计：记录每个连续序列的长度，取最大值
//
// 时间复杂度：O(n) - 每个数字最多被访问两次（一次找起始点，一次在序列中）
// 空间复杂度：O(n) - 哈希表存储所有数字
func LongestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	ans := 0
	for k := range set {
		if set[k-1] {
			continue
		}
		cur := 0
		for set[k] {
			cur++
			k++
		}
		ans = max(ans, cur)
	}
	return ans
}
