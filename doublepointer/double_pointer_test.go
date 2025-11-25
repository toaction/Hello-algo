package doublepointer

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "standard example from LeetCode",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "monotonically increasing heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "all heights are the same",
			height: []int{3, 3, 3, 3, 3},
			want:   12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "standard example from LeetCode",
			args: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "no zeros in array",
			args: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "all zeros in array",
			args: []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.args))
			copy(input, tt.args)
			MoveZeroes(input)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", input, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "standard example from LeetCode",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "no solution",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "multiple duplicate zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "standard example from LeetCode",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "monotonically increasing no water trapped",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "valley shape maximum water trapped",
			height: []int{5, 0, 5},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
