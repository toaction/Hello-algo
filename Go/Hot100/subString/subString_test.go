package subString

import (
	"reflect"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, -1, 0}, 0, 3},
	}

	for _, tt := range tests {
		result := SubarraySum(tt.nums, tt.k)
		if result != tt.want {
			t.Errorf("SubarraySum(%v, %v) = %v, expected %v", tt.nums, tt.k, result, tt.want)
		}
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, tt := range tests {
		result := MinWindow(tt.s, tt.t)
		if result != tt.want {
			t.Errorf("MinWindow(%q, %q) = %q, expected %q", tt.s, tt.t, result, tt.want)
		}
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
	}

	for _, tt := range tests {
		result := MaxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("MaxSlidingWindow(%v, %v) = %v, expected %v", tt.nums, tt.k, result, tt.want)
		}
	}
}
