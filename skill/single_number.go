package skill

/*
136. 只出现一次的数字: https://leetcode.cn/problems/single-number/

给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。

示例 1 ：
	输入：nums = [2,2,1]
	输出：1
*/

func SingleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
