package skill

/*
169. 多数元素: https://leetcode.cn/problems/majority-element/

给定一个大小为 n 的数组 nums ，返回其中的多数元素。
多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

示例 1：
	输入：nums = [3,2,3]
	输出：3
*/

// 投票，众数+1，非众数-1
// 一般情况下，非众数被众数一比一抵消
// 边界情况，非众数之间互相抵消
func MajorityElement(nums []int) int {
	ans, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			ans = v
			count++
			continue
		}
		if v == ans {
			count++
		} else {
			count--
		}
	}
	return ans
}
