package backtrack

/*
131. 分割回文串：https://leetcode.cn/problems/palindrome-partitioning/

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

示例 1：
	输入：s = "aab"
	输出：[["a","a","b"],["aa","b"]]
*/

func Partition(s string) [][]string {
	n := len(s)
	// s[i:j+1] 是否为回文串
	palindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		palindrome[i] = make([]bool, n)
		palindrome[i][i] = true
	}
	// dp 初始化
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				continue
			}
			if j-i < 3 {
				palindrome[i][j] = true
				continue
			}
			palindrome[i][j] = palindrome[i+1][j-1]
		}
	}
	// 回溯找到分割方案
	ans := [][]string{}
	path := []string{}
	backTrackPartition(s, 0, palindrome, &path, &ans)
	return ans
}

func backTrackPartition(s string, i int, palindrome [][]bool, path *[]string, ans *[][]string) {
	if i == len(s) {
		*ans = append(*ans, append([]string{}, *path...))
		return
	}
	for j := i; j < len(s); j++ {
		if !palindrome[i][j] {
			continue
		}
		*path = append(*path, s[i:j+1])
		backTrackPartition(s, j+1, palindrome, path, ans)
		*path = (*path)[:len(*path)-1]
	}
}
