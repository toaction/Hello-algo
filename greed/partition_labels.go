package greed

// 763. 划分字母区间
//
// https://leetcode.cn/problems/partition-labels/
//
// 给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
//
// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//
// 示例 1：
//
//	输入：s = "ababcbacadefegdehijhklij"
//	输出：[9,7,8]
//	解释：
//	    划分结果为 "ababcbaca"、"defegde"、"hijhklij"
//	    每个字母最多出现在一个片段中
//	    像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少
func PartitionLabels(s string) []int {
	ans := make([]int, 0)
	lastIdx := [128]int{}
	// 每个字符最后出现的位置
	for i := 0; i < len(s); i++ {
		lastIdx[s[i]] = i
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 当前区间内所有字符的最后出现位置的最大值
		end = max(end, lastIdx[s[i]])
		if end == i {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
