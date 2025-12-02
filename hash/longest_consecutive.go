package hash

/* 
128. 最长连续序列: https://leetcode-cn.com/problems/longest-consecutive-sequence/

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

示例 1：
	输入：nums = [100,4,200,1,3,2]
	输出：4
	解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。 
*/


// 使用一个 set 存放数组中出现过的元素
// 遍历 set 集合中的元素 x（不要遍历数组，数组重复元素可能很多）
// 若集合存在 x-1, 则代表 x 不是连续序列的起点；否则计算以 x 为起点的连续序列的长度
func LongestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	ans := 0
	for x := range set {
		if set[x-1] {
			continue
		}
		cur := 0
		for set[x] {
			cur++
			x++
		}
		ans = max(ans, cur)
	}
	return ans
}
