package matrix

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "standard leetcode example",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "4x4 matrix",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				input[i] = make([]int, len(tt.matrix[i]))
				copy(input[i], tt.matrix[i])
			}
			Rotate(input)
			if !equalMatrix(input, tt.want) {
				t.Errorf("Rotate() = %v, want %v", input, tt.want)
			}
		})
	}
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "standard example - target found",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
		{
			name: "target not found",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
		{
			name: "single row matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			target: 3,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "standard example",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "multiple zeros",
			matrix: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{1, 0, 1},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 5},
				{0, 0, 0},
			},
		},
		{
			name: "no zeros",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				input[i] = make([]int, len(tt.matrix[i]))
				copy(input[i], tt.matrix[i])
			}
			SetZeroes(input)
			if !equalMatrix(input, tt.want) {
				t.Errorf("SetZeroes() = %v, want %v", input, tt.want)
			}
		})
	}
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name: "standard example",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "rectangular matrix",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "single row",
			matrix: [][]int{
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpiralOrder(tt.matrix)
			if !equalSlices(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
