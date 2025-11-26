package slidewindow

// 438. 找到字符串中所有字母异位词
//
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
//
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例 1:
//
//	输入: s = "cbaebabacd", p = "abc"
//	输出: [0,6]
//	解释:
//	起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//	起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
// 解题思路：
// 1. 滑动窗口：维护一个大小为len(p)的窗口在s上滑动
// 2. 字符统计：用need数组记录p中每个字符的出现次数
// 3. 差值计数：用diff记录还需要匹配的字符总数
// 4. 窗口操作：
//   - 扩展：右指针向右移动，更新字符计数和diff值
//   - 收缩：当窗口包含所有所需字符时，尝试从左侧收缩
//   - 匹配：当窗口大小等于len(p)且diff=0时，找到异位词
//
// 5. 核心优化：
//   - need[x] > 0：表示还需要x字符
//   - need[x] < 0：表示x字符多余，可以从窗口移除
//   - diff=0：表示窗口包含所有必需字符
//
// 时间复杂度：O(m+n) - m是s的长度，n是p的长度
// 空间复杂度：O(1) - 固定大小的字符数组
func FindAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	as := []byte(s)
	ap := []byte(p)
	m, n := len(as), len(ap)
	if m < n {
		return []int{}
	}
	// 存放所需每个字符的个数
	need := [128]int{}
	for _, v := range ap {
		need[v]++
	}
	// 存放还需字符的总个数
	diff := n
	left, right := 0, 0
	for right < m {
		// 扩大窗口，判断右指针指向字符是否为所需字符
		if need[as[right]] > 0 {
			diff--
		}
		need[as[right]]--
		right++
		// 判断窗口是否满足所需字符
		if diff == 0 {
			// 尝试缩小左窗口
			for need[as[left]] < 0 {
				need[as[left]]++
				left++
			}
			// 查看窗口大小是否符合
			if right-left == n {
				ans = append(ans, left)
			}
			// 移动左指针,查找下一个满足条件的窗口
			need[as[left]]++
			left++
			diff++
		}
	}
	return ans
}
