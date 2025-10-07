package subString

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
//
// 示例 1：
//
// 输入：nums = [1,1,1], k = 2
//
// 输出：2
//
// 前缀和 + Hash
func SubarraySum(nums []int, k int) int {
	ans := 0
	// key = 前缀和，value = 个数
	hashMap := make(map[int]int)
	sum := 0
	hashMap[0] = 1
	for _, v := range nums {
		sum += v
		// [i, cur] = k
		ans += hashMap[sum-k]
		hashMap[sum]++
	}
	return ans
}
