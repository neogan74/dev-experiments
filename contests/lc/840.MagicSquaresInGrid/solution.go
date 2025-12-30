package magicsquaresingrid

func isMagicSquare(grid [][]int, row, col int) bool {
	// Extract the 3x3 subgrid
	square := make([][]int, 3)
	for i := 0; i < 3; i++ {
		square[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			square[i][j] = grid[row+i][col+j]
		}
	}

	// Check if all numbers from 1 to 9 are present
	seen := make([]bool, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := square[i][j]
			if num < 1 || num > 9 {
				return false
			}
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check rows
	rowSums := make([]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rowSums[i] += square[i][j]
		}
	}

	// Check if all rows have the same sum
	targetSum := rowSums[0]
	for i := 1; i < 3; i++ {
		if rowSums[i] != targetSum {
			return false
		}
	}

	// Check columns
	colSums := make([]int, 3)
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			colSums[j] += square[i][j]
		}
		if colSums[j] != targetSum {
			return false
		}
	}

	// Check diagonals
	diag1Sum := square[0][0] + square[1][1] + square[2][2]
	diag2Sum := square[0][2] + square[1][1] + square[2][0]
	return diag1Sum == targetSum && diag2Sum == targetSum
}

func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Iterate through all possible top-left positions for a 3x3 subgrid
	for row := 0; row <= rows-3; row++ {
		for col := 0; col <= cols-3; col++ {
			if isMagicSquare(grid, row, col) {
				count++
			}
		}
	}

	return count
}
