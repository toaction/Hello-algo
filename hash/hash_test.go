package hash

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic example",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "duplicate numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "basic example",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "single character strings",
			strs: []string{"a", "b", "c", "a"},
			want: [][]string{{"a", "a"}, {"b"}, {"c"}},
		},
		{
			name: "empty strings",
			strs: []string{"", ""},
			want: [][]string{{"", ""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.strs)

			// Sort each group in result and expected
			for i := range got {
				sort.Strings(got[i])
			}
			for i := range tt.want {
				sort.Strings(tt.want[i])
			}

			// Sort groups by first element for comparison
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) == 0 && len(got[j]) == 0 {
					return false
				}
				if len(got[i]) == 0 {
					return true
				}
				if len(got[j]) == 0 {
					return false
				}
				return got[i][0] < got[j][0]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if len(tt.want[i]) == 0 && len(tt.want[j]) == 0 {
					return false
				}
				if len(tt.want[i]) == 0 {
					return true
				}
				if len(tt.want[j]) == 0 {
					return false
				}
				return tt.want[i][0] < tt.want[j][0]
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "all consecutive",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "duplicates",
			nums: []int{1, 2, 2, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
