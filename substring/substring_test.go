package substring

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "standard example from LeetCode",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "single element window",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all negative numbers",
			nums: []int{-1, -3, -2, -5, -6},
			k:    2,
			want: []int{-1, -2, -2, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "standard example from LeetCode",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "no solution",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "multiple possible windows",
			s:    "abbbbbcd",
			t:    "bcd",
			want: "bcd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "standard example from LeetCode",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "negative numbers",
			nums: []int{1, -1, 0},
			k:    0,
			want: 3, // [0], [1,-1], [1,-1,0]
		},
		{
			name: "large numbers",
			nums: []int{100, 1, 200, 2, 300},
			k:    502,
			want: 1, // [100,1,200,2,300]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
