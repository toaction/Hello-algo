package doublePtr

import (
	"reflect"
	"sort"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
	}

	for _, tt := range tests {
		result := MaxArea(tt.height)
		if result != tt.want {
			t.Errorf("MaxArea(%v) = %v, expected %v", tt.height, result, tt.want)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{0, 0, 0, 1}, []int{1, 0, 0, 0}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		MoveZeroes(nums)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("MoveZeroes(%v) = %v, expected %v", tt.nums, nums, tt.want)
		}
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		result := ThreeSum(tt.nums)
		if !isEqualThreeSum(result, tt.want) {
			t.Errorf("ThreeSum(%v) = %v, expected %v", tt.nums, result, tt.want)
		}
	}
}

func isEqualThreeSum(result, want [][]int) bool {
	if len(result) != len(want) {
		return false
	}

	sortResults(result)
	sortResults(want)

	return reflect.DeepEqual(result, want)
}

func sortResults(res [][]int) {
	for i := range res {
		sort.Ints(res[i])
	}
	sort.Slice(res, func(i, j int) bool {
		for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return len(res[i]) < len(res[j])
	})
}

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{3, 0, 2, 0, 4}, 7},
	}

	for _, tt := range tests {
		result := Trap(tt.height)
		if result != tt.want {
			t.Errorf("Trap(%v) = %v, expected %v", tt.height, result, tt.want)
		}
	}
}
