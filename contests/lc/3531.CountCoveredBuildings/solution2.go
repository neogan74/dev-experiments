package countcoveredbuildings

func countCoveredBuildings2(n int, buildings [][]int) int {
	xmin := make([]int, n+1)
	xmax := make([]int, n+1)
	ymin := make([]int, n+1)
	ymax := make([]int, n+1)
	for i := 0; i <= n; i++ {
		xmin[i] = n + 1
		ymin[i] = n + 1
	}
	for _, b := range buildings {
		xmin[b[0]] = min(xmin[b[0]], b[1])
		xmax[b[0]] = max(xmax[b[0]], b[1])
		ymin[b[1]] = min(ymin[b[1]], b[0])
		ymax[b[1]] = max(ymax[b[1]], b[0])
	}
	var res int
	for _, b := range buildings {
		if 0 < xmin[b[0]] && xmax[b[0]] <= n && 0 < ymin[b[1]] && ymax[b[1]] <= n {
			if xmin[b[0]] < b[1] && b[1] < xmax[b[0]] && ymin[b[1]] < b[0] && b[0] < ymax[b[1]] {
				res++
			}
		}
	}
	return res
}
