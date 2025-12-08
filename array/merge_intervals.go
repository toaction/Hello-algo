package array

import "sort"

/*
56. 合并区间：https://leetcode.cn/problems/merge-intervals/

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
	输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
	输出：[[1,6],[8,10],[15,18]]
	解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func Merge(intervals [][]int) [][]int {
	// 处理空数组的情况
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 先按照 data[0] 排序，然后按照 data[1] 排序
	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	
	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > ans[len(ans)-1][1] {
			// 区间不重叠
			ans = append(ans, intervals[i])
		} else {
			// 区间重叠，合并
			if intervals[i][1] > ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		}
	}
	return ans
}
