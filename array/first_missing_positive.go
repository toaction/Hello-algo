package array

/*
41. 缺失的第一个正数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

示例 1:
	输入：nums = [1,2,0]
	输出：3
	解释：范围 [1,2] 中的数字都在数组中。
*/

// 原地哈希
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	// 将数组中非 [1,n] 的数都置为 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] > n {
			nums[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		k := nums[i] % (n + 1)
		if k == 0 {
			continue
		}
		// 标记正整数 k 存在数组中
		nums[k-1] = nums[k-1] + (n + 1)
	}
	for i, v := range nums {
		if v < n+1 {
			return i + 1
		}
	}
	return n + 1
}
