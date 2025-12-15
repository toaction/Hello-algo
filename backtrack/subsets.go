package backtrack

/*
78. 子集：https://leetcode.cn/problems/subsets/

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
	输入：nums = [1,2,3]
	输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/

func Subsets(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	ans = append(ans, []int{})
	backTrackSubsets(nums, 0, &ans, &path)
	return ans
}

func backTrackSubsets(nums []int, idx int, ans *[][]int, path *[]int) {
	if idx > len(nums) {
		return
	}
	for i := idx; i < len(nums); i++ {
		*path = append(*path, nums[i])
		*ans = append(*ans, append([]int{}, *path...))
		backTrackSubsets(nums, i+1, ans, path)
		*path = (*path)[:len(*path)-1]
	}
}
