package countcoveredbuildings

// countCoveredBuildings returns the number of buildings that have at least
// one building in each cardinal direction (left, right, above, below).
func countCoveredBuildings(n int, buildings [][]int) int {
	// Track smallest and largest column per row, and row per column.
	rowLimits := make(map[int][2]int)
	colLimits := make(map[int][2]int)

	for _, b := range buildings {
		r, c := b[0], b[1]

		if lim, ok := rowLimits[r]; ok {
			rowLimits[r] = [2]int{min(lim[0], c), max(lim[1], c)}
		} else {
			rowLimits[r] = [2]int{c, c}
		}

		if lim, ok := colLimits[c]; ok {
			colLimits[c] = [2]int{min(lim[0], r), max(lim[1], r)}
		} else {
			colLimits[c] = [2]int{r, r}
		}
	}

	covered := 0
	for _, b := range buildings {
		r, c := b[0], b[1]
		rowMin, rowMax := rowLimits[r][0], rowLimits[r][1]
		colMin, colMax := colLimits[c][0], colLimits[c][1]

		if c != rowMin && c != rowMax && r != colMin && r != colMax {
			covered++
		}
	}

	return covered
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
