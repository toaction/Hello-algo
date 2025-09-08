package com.action.hash;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GroupAnagrams {

    /*
    source: https://leetcode.com/problems/group-anagrams
    desc:
        给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
    * */

    public static List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        // 字母异位词使用同一个key：排序、字符拼接
        for(String s:strs){
            String key = getKey(s);
            List<String> value = map.getOrDefault(key, new ArrayList<>());
            value.add(s);
            map.put(key, value);
        }
        return new ArrayList<>(map.values());
    }

    private static String getKey(String s){
        int[] cnt = new int[26];
        int len = s.length();
        for(int i=0; i<len; i++){
            cnt[s.charAt(i)-'a']++;
        }
        StringBuilder sb = new StringBuilder();
        for(int i=0; i<26; i++){
            if(cnt[i] == 0) continue;
            sb.append(cnt[i]);
            sb.append((char)('a'+i));
        }
        return sb.toString();
    }


    public static void main(String[] args) {
        String[] strs = new String[]{"eat","tea","tan","ate","nat","bat"};
        List<List<String>> result = groupAnagrams(strs);
        System.out.println(result);
    }
}
