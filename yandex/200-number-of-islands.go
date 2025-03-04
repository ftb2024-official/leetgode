/*
Given an `M x N` 2D binary grid `grid` which represents a map of `1`s (land) and `0`s (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of
the grid are all surrounded by water.

Example 1:
	Input: grid = [
		["1","1","1","1","0"],
		["1","1","0","1","0"],
		["1","1","0","0","0"],
		["0","0","0","0","0"]
	]
	Output: 1

Example 2:
	Input: grid = [
		["1","1","0","0","0"],
		["1","1","0","0","0"],
		["0","0","1","0","0"],
		["0","0","0","1","1"]
	]
	Output: 3
*/

package yandex

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	var islandCount int
	for row := range rows {
		for col := range cols {
			if grid[row][col] == '1' {
				dfs(grid, row, col)
				// bfs(grid, row, col)
				islandCount++
			}
		}
	}

	return islandCount
}

// DFS approach
func dfs(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0' // make the current land as visited

	// visiting all 4 neighbors
	dfs(grid, row-1, col) // UP
	dfs(grid, row+1, col) // DOWN
	dfs(grid, row, col+1) // RIGHT
	dfs(grid, row, col-1) //LEFT
}

// BFS approach
func bfs(grid [][]byte, row, col int) {
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} // UP, DOWN, RIGHT, LEFT
	queue := [][]int{{row, col}}
	grid[row][col] = '0' // mark as visited

	for len(queue) > 0 {
		// dequeue
		current := queue[0]
		queue = queue[1:]

		// explore all directions
		for _, d := range directions {
			newRow, newCol := current[0]+d[0], current[1]+d[1]
			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == '1' {
				grid[newRow][newCol] = '0'
				queue = append(queue, []int{newRow, newCol}) // enqueue
			}
		}
	}
}
