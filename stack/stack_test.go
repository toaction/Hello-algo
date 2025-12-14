package stack

import (
	"reflect"
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("Push and GetMin operations", func(t *testing.T) {
		ms := Constructor()
		ms.Push(-2)
		ms.Push(0)
		ms.Push(-3)
		if got := ms.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}
	})

	t.Run("Pop and Top operations", func(t *testing.T) {
		ms := Constructor()
		ms.Push(-2)
		ms.Push(0)
		ms.Push(-3)
		ms.Pop()
		if got := ms.Top(); got != 0 {
			t.Errorf("Top() = %v, want 0", got)
		}
		if got := ms.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("Single element operations", func(t *testing.T) {
		ms := Constructor()
		ms.Push(5)
		if got := ms.Top(); got != 5 {
			t.Errorf("Top() = %v, want 5", got)
		}
		if got := ms.GetMin(); got != 5 {
			t.Errorf("GetMin() = %v, want 5", got)
		}
		ms.Pop()
		if len(ms.stack) != 0 {
			t.Errorf("Stack should be empty after pop")
		}
	})
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "simple valid parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "nested valid parentheses",
			s:    "({[]})",
			want: true,
		},
		{
			name: "invalid parentheses",
			s:    "(]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "basic example",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "increasing temperatures",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "all decreasing temperatures",
			temperatures: []int{90, 80, 70, 60},
			want:         []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "simple encoded string",
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			name: "nested encoded string",
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			name: "complex nested string",
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeString(tt.s)
			if got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "basic example",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "all equal heights",
			heights: []int{2, 2, 2, 2},
			want:    8,
		},
		{
			name:    "increasing heights",
			heights: []int{1, 2, 3, 4, 5},
			want:    9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestRectangleArea(tt.heights)
			if got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
