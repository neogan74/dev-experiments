package solution

// CountNegatives counts negative numbers in a row- and column-sorted matrix.
func countNegatives(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	r := rows - 1
	c := 0
	count := 0

	for r >= 0 && c < cols {
		if grid[r][c] < 0 {
			count += cols - c
			r--
		} else {
			c++
		}
	}

	return count
}
