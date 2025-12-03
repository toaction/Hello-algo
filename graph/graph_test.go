package graph

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "single island",
			grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "multiple islands",
			grid: [][]byte{
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
				{'1', '0', '1', '0'},
			},
			want: 6,
		},
		{
			name: "no islands",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of grid for each test to avoid modification
			gridCopy := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]byte, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			got := NumIslands(gridCopy)
			if got != tt.want {
				t.Errorf("NumIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "standard rotting",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "already all rotten",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			want: 0,
		},
		{
			name: "impossible to rot all",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of grid for each test to avoid modification
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			got := OrangesRotting(gridCopy)
			if got != tt.want {
				t.Errorf("OrangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:       "possible schedule",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			want: true,
		},
		{
			name:       "impossible cycle",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			want: false,
		},
		{
			name:          "no prerequisites",
			numCourses:    4,
			prerequisites: [][]int{},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("CanFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	tests := []struct {
		name  string
		words []string
	}{
		{
			name:  "single word",
			words: []string{"hello"},
		},
		{
			name:  "multiple words",
			words: []string{"apple", "app", "application"},
		},
		{
			name:  "words with common prefixes",
			words: []string{"car", "cat", "cart"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, word := range tt.words {
				trie.Insert(word)
			}
			// Test that all inserted words can be found
			for _, word := range tt.words {
				if !trie.Search(word) {
					t.Errorf("Inserted word %s not found", word)
				}
			}
		})
	}
}

func TestTrie_Search(t *testing.T) {
	tests := []struct {
		name   string
		words  []string // words to insert
		search string   // word to search
		want   bool
	}{
		{
			name:   "search existing word",
			words:  []string{"hello", "world"},
			search: "hello",
			want:   true,
		},
		{
			name:   "search non-existing word",
			words:  []string{"hello", "world"},
			search: "test",
			want:   false,
		},
		{
			name:   "search prefix only",
			words:  []string{"apple"},
			search: "app",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, word := range tt.words {
				trie.Insert(word)
			}
			got := trie.Search(tt.search)
			if got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	tests := []struct {
		name   string
		words  []string // words to insert
		prefix string   // prefix to search
		want   bool
	}{
		{
			name:   "prefix exists",
			words:  []string{"apple", "application"},
			prefix: "app",
			want:   true,
		},
		{
			name:   "prefix does not exist",
			words:  []string{"hello", "world"},
			prefix: "xyz",
			want:   false,
		},
		{
			name:   "full word as prefix",
			words:  []string{"hello"},
			prefix: "hello",
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, word := range tt.words {
				trie.Insert(word)
			}
			got := trie.StartsWith(tt.prefix)
			if got != tt.want {
				t.Errorf("Trie.StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
