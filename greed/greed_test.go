package greed

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "multiple partitions",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name: "single partition",
			s:    "eccbbbbdec",
			want: []int{10},
		},
		{
			name: "each character once",
			s:    "abcdef",
			want: []int{1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionLabels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic case",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "alternative path",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Jump(tt.nums)
			if got != tt.want {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "can reach end",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "blocked by zero",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "single element",
			nums: []int{0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanJump(tt.nums)
			if got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "profit available",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "no profit - decreasing",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "profit at end",
			prices: []int{2, 4, 1, 7},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
