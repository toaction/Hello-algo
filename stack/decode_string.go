package stack

import "strconv"

/*
394. 字符串解码：https://leetcode.cn/problems/decode-string/

给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
	输入：s = "3[a]2[bc]"
	输出："aaabcbc"
*/

func DecodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
			continue
		}
		// 找到编码字符串
		k := len(stack) - 1
		for stack[k] != '[' {
			k--
		}
		str := make([]byte, 0)
		str = append(str, stack[k+1:]...)
		// 找到重复次数
		idx := k - 1
		for idx >= 0 && stack[idx] >= '0' && stack[idx] <= '9' {
			idx--
		}
		num, _ := strconv.Atoi(string(stack[idx+1 : k]))
		// 重新入栈
		stack = stack[:idx+1]
		for j := 0; j < num; j++ {
			stack = append(stack, str...)
		}
	}
	return string(stack)
}
