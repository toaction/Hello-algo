package hash

import "sort"

/* 
49. 字母异位词分组: https://leetcode.cn/problems/group-anagrams/

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:
	输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
	输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
	解释：
	在 strs 中没有字符串可以通过重新排列来形成 "bat"。
	字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
	字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。 
*/


// 使用一个哈希表存储互为字母异位词的字符串列表，
// 互为字母异位词的字符串对应的 key 需相同，有以下思路:
//	1. 对字符串排序，将排序结果作为 key
//	2. 统计字符串中每个字母出现的次数
func GroupAnagramsBaseSort(strs []string) [][]string {
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

func GroupAnagramsBaseCount(strs []string) [][]string {
	// key = count(s), value = 字母异位词集合
	mp := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c - 'a']++
		}
		mp[key] = append(mp[key], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
