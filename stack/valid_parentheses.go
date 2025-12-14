package stack

/*
20. 有效的括号: https://leetcode.cn/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

示例 1：
	输入：s = "()"
	输出：true
*/

// 使用栈存放左括号，遇到右括号弹出栈顶元素，查看是否匹配
func IsValid(s string) bool {
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		// 左括号
		if pairs[s[i]] > 0 {
			stack = append(stack, s[i])
			continue
		}
		// 无左括号 或 左括号不匹配
		if len(stack) == 0 || pairs[stack[len(stack)-1]] != s[i] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
