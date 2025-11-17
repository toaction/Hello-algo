package linkedlist

import (
	"reflect"
	"testing"
)

// createList creates a linked list from a slice of integers
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

// listToSlice converts a linked list to a slice of integers
func listToSlice(head *ListNode) []int {
	result := []int{} // initialize as empty slice instead of nil
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// TestAddTwoNumbers tests the AddTwoNumbers function
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "basic case: 342 + 465 = 807",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "with carry: 9999 + 1 = 10000",
			l1:   []int{9, 9, 9, 9},
			l2:   []int{1},
			want: []int{0, 0, 0, 0, 1},
		},
		{
			name: "simple addition: 0 + 0 = 0",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)
			got := AddTwoNumbers(l1, l2)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestRemoveNthFromEnd tests the RemoveNthFromEnd function
func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "remove 2nd node from end",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "remove only node",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "remove head node (2nd from end)",
			head: []int{1, 2},
			n:    2,
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := RemoveNthFromEnd(head, tt.n)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestHasCycle tests the HasCycle function
func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{
			name: "linked list with cycle",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			name: "linked list without cycle",
			vals: []int{1, 2, 3, 4, 5},
			pos:  -1,
			want: false,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.vals, tt.pos)
			got := HasCycle(head)
			if got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// createCycleList creates a linked list with a cycle at the given position
// pos is the index where the tail connects to. pos = -1 means no cycle.
func createCycleList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	cur := head
	var cycleNode *ListNode

	if pos == 0 {
		cycleNode = head
	}

	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
		if i == pos {
			cycleNode = cur
		}
	}

	if pos >= 0 && cycleNode != nil {
		cur.Next = cycleNode
	}

	return head
}
