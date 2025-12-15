package backtrack

/*
46. 全排列：https://leetcode.cn/problems/permutations/

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func Permute(nums []int) [][]int {
	ans := &[][]int{}
	path := &[]int{}
	used := make([]bool, len(nums))
	backTrackPermute(nums, used, ans, path)
	return *ans
}

func backTrackPermute(nums []int, used []bool, ans *[][]int, path *[]int) {
	if len(*path) == len(nums) {
		*ans = append(*ans, append([]int{}, *path...))
		return
	}
	for i, v := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		*path = append(*path, v)
		backTrackPermute(nums, used, ans, path)
		*path = (*path)[:len(*path)-1]
		used[i] = false
	}
}
