package tree

import (
	"testing"
)

// Helper function to create a binary tree for testing
func createTestTree() *TreeNode {
	//      1
	//     / \
	//    2   3
	//   / \
	//  4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	return root
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{4, 2, 5, 1, 3}, // left, root, right, left subtree, right
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("InorderTraversal() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestInorderTraversalBaseIteration(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{4, 2, 5, 1, 3},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversalBaseIteration(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("InorderTraversalBaseIteration() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("InorderTraversalBaseIteration() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{1, 2, 4, 5, 3}, // root, left subtree, right
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreorderTraversal(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("PreorderTraversal() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestPreorderTraversalBaseInteration(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{1, 2, 4, 5, 3},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreorderTraversalBaseInteration(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("PreorderTraversalBaseInteration() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("PreorderTraversalBaseInteration() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{4, 5, 2, 3, 1}, // left subtree, right, root
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostorderTraversal(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("PostorderTraversal() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestPostorderTraversalBaseInteration(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: []int{4, 5, 2, 3, 1},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "empty tree",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostorderTraversalBaseInteration(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("PostorderTraversalBaseInteration() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("PostorderTraversalBaseInteration() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: [][]int{
				{1},    // level 0: root
				{2, 3}, // level 1: children
				{4, 5}, // level 2: grandchildren
			},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "empty tree",
			root: nil,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrder(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("LevelOrder() levels = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("LevelOrder() level %d length = %v, want %v", i, len(got[i]), len(tt.want[i]))
					continue
				}
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
						break
					}
				}
			}
		})
	}
}

func TestLevelOrderBaseInteration(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "standard binary tree",
			root: createTestTree(),
			want: [][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "empty tree",
			root: nil,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrderBaseInteration(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("LevelOrderBaseInteration() levels = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("LevelOrderBaseInteration() level %d length = %v, want %v", i, len(got[i]), len(tt.want[i]))
					continue
				}
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("LevelOrderBaseInteration() = %v, want %v", got, tt.want)
						break
					}
				}
			}
		})
	}
}
