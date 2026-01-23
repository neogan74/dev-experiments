package main

func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	rowPref := make([][]int, m)
	for i := 0; i < m; i++ {
		rowPref[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			rowPref[i][j+1] = rowPref[i][j] + grid[i][j]
		}
	}

	colPref := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		colPref[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			colPref[i+1][j] = colPref[i][j] + grid[i][j]
		}
	}

	diagPref := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		diagPref[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagPref[i+1][j+1] = diagPref[i][j] + grid[i][j]
		}
	}

	antiPref := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		antiPref[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			antiPref[i+1][j] = antiPref[i][j+1] + grid[i][j]
		}
	}

	maxSize := m
	if n < maxSize {
		maxSize = n
	}
	for k := maxSize; k >= 2; k-- {
		for i := 0; i+k <= m; i++ {
			for j := 0; j+k <= n; j++ {
				target := rowPref[i][j+k] - rowPref[i][j]
				diag := diagPref[i+k][j+k] - diagPref[i][j]
				if diag != target {
					continue
				}
				anti := antiPref[i+k][j] - antiPref[i][j+k]
				if anti != target {
					continue
				}
				ok := true
				for r := 0; r < k && ok; r++ {
					rowSum := rowPref[i+r][j+k] - rowPref[i+r][j]
					if rowSum != target {
						ok = false
						break
					}
					colSum := colPref[i+k][j+r] - colPref[i][j+r]
					if colSum != target {
						ok = false
						break
					}
				}
				if ok {
					return k
				}
			}
		}
	}
	return 1
}

func main() {}
