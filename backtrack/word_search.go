package backtrack

/*
79. 单词搜索

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

示例 1：
	输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	输出：true
*/

func WordExist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] != word[0] {
				continue
			}
			visited[i][j] = true
			if backTrackExist(board, word, visited, 1, i, j) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func backTrackExist(board [][]byte, word string, visited [][]bool, idx int, prevI, prevJ int) bool {
	if idx == len(word) {
		return true
	}
	h, w := len(board), len(board[0])
	for _, v := range directions {
		i, j := prevI+v[0], prevJ+v[1]
		if i >= 0 && i < h && j >= 0 && j < w && !visited[i][j] && board[i][j] == word[idx] {
			visited[i][j] = true
			if backTrackExist(board, word, visited, idx+1, i, j) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}
