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

// TestDetectCycle tests the DetectCycle function
func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int // position of cycle, -1 means no cycle
		want int // expected cycle node value, -1 means no cycle
	}{
		{
			name: "cycle exists at position 1",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: 2,
		},
		{
			name: "cycle at position 0",
			vals: []int{1, 2},
			pos:  0,
			want: 1,
		},
		{
			name: "no cycle",
			vals: []int{1, 2, 3, 4, 5},
			pos:  -1,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.vals, tt.pos)
			got := DetectCycle(head)
			if tt.want == -1 {
				if got != nil {
					t.Errorf("DetectCycle() = %v, want nil", got)
				}
			} else {
				if got == nil || got.Val != tt.want {
					t.Errorf("DetectCycle() = %v, want node with value %d", got, tt.want)
				}
			}
		})
	}
}

// TestGetIntersectionNode tests the GetIntersectionNode function
func TestGetIntersectionNode(t *testing.T) {
	// Helper function to create intersecting lists
	createIntersectingLists := func(aVals, bVals []int, intersectPos int) (*ListNode, *ListNode) {
		headA := createList(aVals)
		headB := createList(bVals)

		if intersectPos >= 0 {
			// Find intersection node in list A
			intersectNode := headA
			for i := 0; i < intersectPos; i++ {
				intersectNode = intersectNode.Next
			}

			// Find tail of list B
			tailB := headB
			for tailB.Next != nil {
				tailB = tailB.Next
			}

			// Connect tail of B to intersection node
			tailB.Next = intersectNode
		}

		return headA, headB
	}

	tests := []struct {
		name       string
		aVals      []int
		bVals      []int
		intersectPos int // position in list A where intersection occurs
		want       int   // expected intersection value, -1 means no intersection
	}{
		{
			name: "intersect at position 2",
			aVals: []int{4, 1, 8, 4, 5},
			bVals: []int{5, 6, 1},
			intersectPos: 2,
			want: 8,
		},
		{
			name: "no intersection",
			aVals: []int{2, 6, 4},
			bVals: []int{1, 5},
			intersectPos: -1,
			want: -1,
		},
		{
			name: "intersect at head",
			aVals: []int{1, 2, 3},
			bVals: []int{9, 8},
			intersectPos: 0,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB := createIntersectingLists(tt.aVals, tt.bVals, tt.intersectPos)
			got := GetIntersectionNode(headA, headB)
			if tt.want == -1 {
				if got != nil {
					t.Errorf("GetIntersectionNode() = %v, want nil", got)
				}
			} else {
				if got == nil || got.Val != tt.want {
					t.Errorf("GetIntersectionNode() = %v, want node with value %d", got, tt.want)
				}
			}
		})
	}
}

// TestLRUCache tests the LRUCache operations
func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string // "put" or "get"
		keys     []int
		values   []int // for put operations
		capacity int
		expected []int // for get operations, -1 means check no error
	}{
		{
			name:     "basic operations",
			ops:      []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			keys:     []int{1, 2, 1, 3, 2, 4, 1, 3, 4},
			values:   []int{1, 2, 0, 3, 0, 4, 0, 0, 0},
			capacity: 2,
			expected: []int{-1, -1, 1, -1, -1, -1, -1, 3, 4},
		},
		{
			name:     "single capacity",
			ops:      []string{"put", "put", "get", "put", "get"},
			keys:     []int{1, 2, 1, 3, 2},
			values:   []int{1, 2, 0, 3, 0},
			capacity: 1,
			expected: []int{-1, -1, -1, -1, -1},
		},
		{
			name:     "negative test - non-existent key",
			ops:      []string{"get", "put", "get", "get"},
			keys:     []int{1, 2, 2, 3},
			values:   []int{0, 1, 0, 0},
			capacity: 3,
			expected: []int{-1, -1, 1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			for i, op := range tt.ops {
				if op == "put" {
					cache.Put(tt.keys[i], tt.values[i])
				} else if op == "get" {
					got := cache.Get(tt.keys[i])
					want := tt.expected[i]
					if want != -1 && got != want {
						t.Errorf("LRUCache.Get(%d) = %d, want %d", tt.keys[i], got, want)
					}
				}
			}
		})
	}
}

// TestMergeTwoLists tests the MergeTwoLists function
func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "both non-empty lists",
			l1:   []int{1, 2, 4},
			l2:   []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "one empty list",
			l1:   []int{},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "both empty lists",
			l1:   []int{},
			l2:   []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)
			got := MergeTwoLists(l1, l2)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestMergeKLists tests the MergeKLists function
func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "multiple lists",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "empty input",
			lists: [][]int{},
			want:  []int{},
		},
		{
			name:  "contains empty lists",
			lists: [][]int{{}, {1}, {2, 3}, {}},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			for _, vals := range tt.lists {
				lists = append(lists, createList(vals))
			}
			got := MergeKLists(lists)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("MergeKLists() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestReverseList tests the ReverseList function
func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "regular list",
			head: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "single node",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "empty list",
			head: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := ReverseList(head)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestSortList tests the SortList function
func TestSortList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "unsorted list",
			head: []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "already sorted",
			head: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "single node",
			head: []int{5},
			want: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := SortList(head)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("SortList() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestSwapPairs tests the SwapPairs function
func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "even number of nodes",
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			name: "odd number of nodes",
			head: []int{1, 2, 3},
			want: []int{2, 1, 3},
		},
		{
			name: "single node",
			head: []int{1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := SwapPairs(head)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("SwapPairs() = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestIsPalindrome tests the IsPalindrome function
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want bool
	}{
		{
			name: "palindrome list",
			head: []int{1, 2, 2, 1},
			want: true,
		},
		{
			name: "non-palindrome list",
			head: []int{1, 2},
			want: false,
		},
		{
			name: "single node",
			head: []int{1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := IsPalindrome(head)
			if got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestReverseKGroup tests the ReverseKGroup function
func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			name: "k=2, complete groups",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "k=3, incomplete last group",
			head: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "k=1, no change",
			head: []int{1, 2, 3},
			k:    1,
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.head)
			got := ReverseKGroup(head, tt.k)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReverseKGroup() = %v, want %v", result, tt.want)
			}
		})
	}
}
