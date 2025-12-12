package substring

/* 
560. 和为 k 的子数组：https://leetcode.cn/problems/subarray-sum-equals-k/

给定一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：
	输入：nums = [1,1,1], k = 2
	输出：2 
*/


// 前缀和 + 哈希表
func SubarraySum(nums []int, k int) int {
    // sum: 从数组开头到当前位置的累计和
    ans, sum := 0, 0
    // 前缀和：个数 (前缀和不包含 nums[i] 本身)
    count := map[int]int{}
    for _, num := range nums {
        count[sum]++
        sum += num
        // 子数组 nums[i, cur] 的和 == k
        ans += count[sum - k]
    }
    return ans
}
