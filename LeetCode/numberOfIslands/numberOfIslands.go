package numberOfIslands

func numIslands(grid [][]byte) int {
	var cnt int

	for column, rows := range grid {
		for row, v := range rows {
			if v != 1 {
				continue
			}
			cnt++
			numIslandsSub(grid, column, row)
		}
	}

	return cnt
}

func numIslandsSub(grid [][]byte, column, row int) {
	if column < 0 || row < 0 || column >= len(grid) || row >= len(grid[column]) || grid[column][row] != 1 {
		return
	}
	grid[column][row] = 2

	numIslandsSub(grid, column+1, row)
	numIslandsSub(grid, column, row+1)
	numIslandsSub(grid, column-1, row)
	numIslandsSub(grid, column, row-1)
}
