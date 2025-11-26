package hash

import "sort"

// 49. 字母异位词分组
//
// https://leetcode.cn/problems/group-anagrams/
//
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 示例 1:
//
//	输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//	输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//	解释：
//	在 strs 中没有字符串可以通过重新排列来形成 "bat"。
//	字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
//	字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
//
// 解题思路：
//	- 特征提取：字母异位词排序后的字符串完全相同
//	- 哈希分组：key = 排序后的字符串，value = 原始字符串集合
//	- 逐个处理：对每个字符串进行排序，作为哈希表的键
//	- 分类收集：将相同key的字符串归类到同一个组
//
// 变种思路：
//   - 字母计数：统计每个字母出现次数作为key（更优，O(L)时间）
//   - 直接使用：本题字符串长度较小，排序方法简单直接
//
// 时间复杂度：O(n * k * logk) - n个字符串，每个长度为k，排序O(klogk)
// 空间复杂度：O(n * k) - 需要存储所有字符串及其排序结果
func GroupAnagrams(strs []string) [][]string {
	// key = sorted(s), value = 字母异位词集合
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		key := string(s)
		mp[key] = append(mp[key], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
