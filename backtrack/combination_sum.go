package backtrack

import "sort"

/*
39. 组合总和：https://leetcode.cn/problems/combination-sum/

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。
你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。
如果至少一个数字的被选数量不同，则两种组合是不同的。

示例 1：
	输入：candidates = [2,3,6,7], target = 7
	输出：[[2,2,3],[7]]
*/

func CombinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backTrackSum(candidates, target, 0, &ans, &path)
	return ans
}

func backTrackSum(candidates []int, target int, idx int, ans *[][]int, path *[]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, *path...))
		return
	}
	for i := idx; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		*path = append(*path, candidates[i])
		backTrackSum(candidates, target-candidates[i], i, ans, path)
		*path = (*path)[:len(*path)-1]
	}
}
