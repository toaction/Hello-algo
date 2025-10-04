package slideWindow

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例 1:
//
// 输入: s = "cbaebabacd", p = "abc"
//
// 输出: [0,6]
//
// 解释:
//
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
func FindAnagrams(s string, p string) []int {
	res := make([]int, 0)
	as := []byte(s)
	ap := []byte(p)
	m, n := len(as), len(ap)
	if m < n {
		return []int{}
	}
	// 存放所需每个字符个数
	an := [128]int{}
	for _, v := range ap {
		an[v]++
	}
	// 存放还需字符个数
	need := n
	left, right := 0, 0
	for right < m {
		// 扩大窗口，判断右指针指向字符是否为所需字符
		if an[as[right]] > 0 {
			need--
		}
		an[as[right]]--
		right++
		// 判断窗口是否满足所需字符
		if need == 0 {
			// 尝试缩小左窗口
			for an[as[left]] < 0 {
				an[as[left]]++
				left++
			}
			// 查看窗口大小是否符合
			if right-left == n {
				res = append(res, left)
			}
			// 移动左指针,查找下一个满足条件的窗口
			an[as[left]]++
			left++
			need++
		}
	}
	return res
}
