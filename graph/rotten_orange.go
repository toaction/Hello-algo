package graph

/*
994. 腐烂的橘子：https://leetcode.cn/problems/rotting-oranges/

在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
	值 0 代表空单元格；
	值 1 代表新鲜橘子；
	值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

示例 1：
	输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
	输出：4
*/

// 广度优先搜索
// 使用一个队列记录上一轮腐烂的橘子
func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 新鲜橘子的个数
	count := 0
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	// 广度优先搜索
	time := 0
	for count > 0 && len(queue) > 0 {
		// 上一分钟新增的腐烂橘子个数
		size := len(queue)
		for k := 0; k < size; k++ {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			if i+1 < m && grid[i+1][j] == 1 {
				count--
				grid[i+1][j] = 2
				queue = append(queue, []int{i + 1, j})
			}
			if i-1 >= 0 && grid[i-1][j] == 1 {
				count--
				grid[i-1][j] = 2
				queue = append(queue, []int{i - 1, j})
			}
			if j+1 < n && grid[i][j+1] == 1 {
				count--
				grid[i][j+1] = 2
				queue = append(queue, []int{i, j + 1})
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				count--
				grid[i][j-1] = 2
				queue = append(queue, []int{i, j - 1})
			}
		}
		time++
	}
	if count > 0 {
		return -1
	}
	return time
}
