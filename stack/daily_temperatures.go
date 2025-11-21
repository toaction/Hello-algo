package stack

// 739. 每日温度
//
// https://leetcode.cn/problems/daily-temperatures/
//
// 给定一个整数数组 temperatures ，表示每天的温度，
// 返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 示例 1:
//
//	输入: temperatures = [73,74,75,71,69,72,76,73]
//	输出: [1,1,4,2,1,1,0,0]
//
// 思路：单调栈，存放温度递减的元素下表
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			k := stack[len(stack)-1]
			ans[k] = i - k
			stack = stack[:len(stack)-1]
		}
		// 存放元素下标
		stack = append(stack, i)
	}
	return ans
}
