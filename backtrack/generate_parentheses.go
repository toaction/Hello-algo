package backtrack

/*
22. 括号生成：https://leetcode.cn/problems/generate-parentheses/

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
	输入：n = 3
	输出：["((()))","(()())","(())()","()(())","()()()"]
*/

func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)
	backTrackParenthesis(n, 0, 0, &ans, &path)
	return ans
}

func backTrackParenthesis(n int, left int, right int, ans *[]string, path *[]byte) {
	if left == n && right == n {
		*ans = append(*ans, string(*path))
		return
	}
	if left < n {
		*path = append(*path, '(')
		backTrackParenthesis(n, left+1, right, ans, path)
		*path = (*path)[:len(*path)-1]
	}
	// 每一个右括号左边的左括号个数要大于右括号个数
	if right < left {
		*path = append(*path, ')')
		backTrackParenthesis(n, left, right+1, ans, path)
		*path = (*path)[:len(*path)-1]
	}
}
