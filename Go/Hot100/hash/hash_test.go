package hash

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{-2, 2, 4}, 2, []int{0, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		result := TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v, expected %v", tt.nums, tt.target, result, tt.want)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{"a", "b", "c"},
			[][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			[]string{"abc", "bca", "cab"},
			[][]string{{"abc", "bca", "cab"}},
		},
	}

	for _, tt := range tests {
		result := GroupAnagrams(tt.strs)
		if !isEqualGroupAnagrams(result, tt.want) {
			t.Errorf("GroupAnagrams(%v) = %v, expected %v", tt.strs, result, tt.want)
		}
	}
}

func isEqualGroupAnagrams(result, want [][]string) bool {
	if len(result) != len(want) {
		return false
	}

	resultMap := make(map[string][]string)
	wantMap := make(map[string][]string)

	for _, group := range result {
		if len(group) > 0 {
			sort.Strings(group)
			key := group[0]
			resultMap[key] = group
		}
	}

	for _, group := range want {
		if len(group) > 0 {
			sort.Strings(group)
			key := group[0]
			wantMap[key] = group
		}
	}

	for key, group := range resultMap {
		wantGroup, exists := wantMap[key]
		if !exists || !reflect.DeepEqual(group, wantGroup) {
			return false
		}
	}

	return true
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 1, 2}, 2},
	}

	for _, tt := range tests {
		result := LongestConsecutive(tt.nums)
		if result != tt.want {
			t.Errorf("LongestConsecutive(%v) = %v, expected %v", tt.nums, result, tt.want)
		}
	}
}
