package stack

/*
84. 柱状图中最大的矩形：https://leetcode.cn/problems/largest-rectangle-in-histogram/

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
	输入： heights = [2,1,5,6,2,3]
	输出： 10
*/

// 单调栈，存放柱子高度递增的柱子下标
// 通过单调栈找到左边小于当前柱子高度的柱子下标（左边界）
// 以及右边小于当前柱子高度的柱子下表（右边界）
// 以当前柱子为高，计算最大宽度对应的面积
func LargestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// 初始化右边界
	for i := 0; i < n; i++ {
		right[i] = n
	}
	// 遍历元素，得到左边界，更新右边界
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			// 栈中元素对应高度递增，因此右边第一个小于柱子的下标为 i
			right[idx] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(right[i]-left[i]-1))
	}
	return ans
}
