package normalArray

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
// 示例 1：
//
// 输入：nums = [1,2,0]
// 输出：3
// 解释：范围 [1,2] 中的数字都在数组中。
//
// 原地哈希：
// 正整数 [1,n] 可能存在数组中，因此使用 n+1 作为模数
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	// 将数组中非 [1,n] 的数都置为 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] > n {
			nums[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		cur := nums[i] % (n + 1)
		if cur == 0 {
			continue
		}
		// 标记正整数 cur 存在数组中
		nums[cur-1] = nums[cur-1] + (n + 1)
	}
	for i := 0; i < n; i++ {
		if nums[i] < n+1 {
			return i + 1
		}
	}
	return n + 1
}
