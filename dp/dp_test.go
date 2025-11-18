package dp

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "two steps",
			n:    2,
			want: 2,
		},
		{
			name: "three steps",
			n:    3,
			want: 3,
		},
		{
			name: "five steps",
			n:    5,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.n)
			if got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic case",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "alternating high values",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "single house",
			nums: []int{5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob(tt.nums)
			if got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "one row",
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			name:    "three rows",
			numRows: 3,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name:    "five rows",
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.numRows)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
