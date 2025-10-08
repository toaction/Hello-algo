package normalArray

import "sort"

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
// 示例 1：
//
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
// 排序
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := intervals[:1]
	idx := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > ans[idx][1] {
			ans = append(ans, intervals[i])
			idx++
		} else {
			ans[idx][1] = max(ans[idx][1], intervals[i][1])
		}
	}
	return ans
}
